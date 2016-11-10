package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Reduce(f interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.reduceStream(f)
	}

	return g.reduce(f)
}

func (g *gollection) reduce(f interface{}) *gollection {
	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Reduce called with non-slice value of type %T", g.slice),
		}
	}

	if sv.Len() == 0 {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Reduce called with empty slice of type %T", g.slice),
		}
	} else if sv.Len() == 1 {
		return &gollection{
			val: sv.Index(0).Interface(),
		}
	}

	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 2 || funcType.NumOut() != 1 {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Reduce called with invalid func. required func(in1, in2 <T>) out <T> but supplied %v", g.slice),
		}
	}

	ret := sv.Index(0).Interface()
	for i := 1; i < sv.Len(); i++ {
		v1 := reflect.ValueOf(ret)
		v2 := sv.Index(i)
		ret = funcValue.Call([]reflect.Value{v1, v2})[0].Interface()
	}

	return &gollection{
		val: ret,
	}
}

func (g *gollection) reduceStream(f interface{}) *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 2 || funcType.NumOut() != 1 {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Reduce called with invalid func. required func(in1, in2 <T>) out <T> but supplied %v", g.slice),
		}
	}

	go func() {
		var ret interface{}
		var initialized bool

		for {
			select {
			case v, ok := <-g.ch:
				if !initialized {
					ret = reflect.ValueOf(v).Interface()
					initialized = true
					continue
				}

				if ok {
					v1 := reflect.ValueOf(ret)
					v2 := reflect.ValueOf(v)
					ret = funcValue.Call([]reflect.Value{v1, v2})[0].Interface()
				} else {
					next.ch <- ret
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
