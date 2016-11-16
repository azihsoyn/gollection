package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Filter(f interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.filterStream(f)
	}
	return g.filter(f)
}

func (g *gollection) validateFilterFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 || funcType.Out(0).Kind() != reflect.Bool {
		return reflect.Value{}, nil, fmt.Errorf("gollection.Filter called with invalid func. required func(in <T>) bool but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) validateSlice(funcName string) (reflect.Value, error) {
	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return reflect.Value{}, fmt.Errorf("gollection.%s called with non-slice value of type %T", funcName, g.slice)
	}
	return sv, nil
}

func (g *gollection) filter(f interface{}) *gollection {
	sv, err := g.validateSlice("Filter")
	if err != nil {
		return &gollection{
			err: err,
		}
	}

	funcValue, funcType, err := g.validateFilterFunc(f)
	if err != nil {
		return &gollection{
			err: err,
		}
	}

	resultSliceType := reflect.SliceOf(funcType.In(0))
	ret := reflect.MakeSlice(resultSliceType, 0, sv.Len())

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if funcValue.Call([]reflect.Value{v})[0].Interface().(bool) {
			ret = reflect.Append(ret, v)
		}
	}

	return &gollection{
		slice: ret.Interface(),
	}
}

func (g *gollection) filterStream(f interface{}) *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	funcValue, funcType, err := g.validateFilterFunc(f)
	if err != nil {
		return &gollection{
			err: err,
		}
	}

	var initialized bool
	go func() {
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					if !initialized {
						// initialize next stream type
						next.ch <- reflect.SliceOf(funcType.In(0))
						initialized = true
						continue
					}

					if funcValue.Call([]reflect.Value{reflect.ValueOf(v)})[0].Interface().(bool) {
						next.ch <- v
					}
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
