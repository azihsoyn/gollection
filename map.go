package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Map(f interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.mapStream(f)
	}

	return g.map_(f)
}

func (g *gollection) map_(f interface{}) *gollection {
	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Map called with non-slice value of type %T", g.slice),
		}
	}

	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Map called with invalid func. required func(in <T>) out <T> but supplied %v", g.slice),
		}
	}
	resultSliceType := reflect.SliceOf(funcType.Out(0))
	ret := reflect.MakeSlice(resultSliceType, 0, sv.Len())

	// avoid "panic: reflect: call of reflect.Value.Interface on zero Value"
	// see https://github.com/azihsoyn/gollection/issues/7
	if sv.Len() == 0 {
		return &gollection{
			slice: ret.Interface(),
		}
	}

	for i := 0; i < sv.Len(); i++ {
		v := funcValue.Call([]reflect.Value{sv.Index(i)})[0]
		ret = reflect.Append(ret, v)
	}

	return &gollection{
		slice: ret.Interface(),
	}

}

func (g *gollection) mapStream(f interface{}) *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 {
		return &gollection{
			err: fmt.Errorf("gollection.Map called with invalid func. required func(in <T>) out <T> but supplied %v", f),
		}
	}

	var initialized bool
	go func() {
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					// initialize next stream type
					if !initialized {
						next.ch <- reflect.SliceOf(funcType.Out(0))
						initialized = true
						continue
					}

					v := funcValue.Call([]reflect.Value{reflect.ValueOf(v)})[0].Interface()
					next.ch <- v
				} else {
					close(next.ch)
					return
				}
			default:
				continue
			}
		}
	}()
	return next
}
