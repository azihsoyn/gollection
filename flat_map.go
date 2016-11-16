package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) FlatMap(f interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.flatMapStream(f)
	}

	return g.flatMap(f)
}

func (g *gollection) validateFlatMapFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 {
		return reflect.Value{}, nil, fmt.Errorf("gollection.FlatMap called with invalid func. required func(in <T>) out <T> but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) flatMap(f interface{}) *gollection {
	sv, err := g.validateSlice("FlatMap")
	if err != nil {
		return &gollection{err: err}
	}

	if err := g.validateSliceOfSlice("FlatMap"); err != nil {
		return &gollection{err: err}
	}

	funcValue, funcType, err := g.validateFlatMapFunc(f)
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
		v := sv.Index(i).Interface()
		svv := reflect.ValueOf(v)
		for j := 0; j < svv.Len(); j++ {
			v := funcValue.Call([]reflect.Value{svv.Index(j)})[0]
			ret = reflect.Append(ret, v)
		}
	}

	return &gollection{
		slice: ret.Interface(),
		err:   nil,
	}
}

func (g *gollection) flatMapStream(f interface{}) *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	funcValue, funcType, err := g.validateFlatMapFunc(f)
	if err != nil {
		return &gollection{err: err}
	}

	var initialized bool
	go func() {
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					if !initialized {
						// initialze next stream type
						currentType := v.(reflect.Type).Elem()
						if currentType.Kind() != reflect.Slice {
							next.ch <- fmt.Errorf("gollection.FlatMap called with non-slice-of-slice value of type %s", currentType)
						}
						next.ch <- reflect.SliceOf(funcType.Out(0))
						initialized = true
						continue
					}

					svv := reflect.ValueOf(v)
					for j := 0; j < svv.Len(); j++ {
						v := funcValue.Call([]reflect.Value{svv.Index(j)})[0]
						next.ch <- v.Interface()
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
