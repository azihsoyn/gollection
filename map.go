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

func (g *gollection) validateMapFunc(f /* func(v <T1>) <T2> */ interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 1 ||
		funcType.NumOut() != 1 {
		return reflect.Value{}, nil, fmt.Errorf("gollection.Map called with invalid func. required func(in <T>) out <T> but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) map_(f interface{}) *gollection {
	sv, err := g.validateSlice("Map")
	if err != nil {
		return &gollection{err: err}
	}

	funcValue, funcType, err := g.validateMapFunc(f)
	if err != nil {
		return &gollection{err: err}
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

	funcValue, funcType, err := g.validateMapFunc(f)
	if err != nil {
		return &gollection{err: err}
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
