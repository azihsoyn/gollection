package gollection

import (
	"fmt"
	"reflect"
	"sync"
)

func (g *gollection) Fold(v0 interface{}, f interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.foldStream(v0, f)
	}

	return g.fold(v0, f)
}

func (g *gollection) validateFoldFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 2 ||
		funcType.NumOut() != 1 {
		return reflect.Value{}, nil, fmt.Errorf("gollection.Fold called with invalid func. required func(in1, in2 <T>) out <T> but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) fold(v0 interface{}, f interface{}) *gollection {
	sv, err := g.validateSlice("Fold")
	if err != nil {
		return &gollection{err: err}
	}

	if sv.Len() < 1 {
		return &gollection{val: v0}
	}

	funcValue, _, err := g.validateFoldFunc(f)
	if err != nil {
		return &gollection{err: err}
	}

	ret := v0
	for i := 0; i < sv.Len(); i++ {
		v1 := reflect.ValueOf(ret)
		v2 := sv.Index(i)
		ret = funcValue.Call([]reflect.Value{v1, v2})[0].Interface()
	}

	return &gollection{
		val: ret,
	}

}

func (g *gollection) foldStream(v0 interface{}, f interface{}) *gollection {
	funcValue, _, err := g.validateFoldFunc(f)
	if err != nil {
		return &gollection{err: err}
	}

	var ret interface{}
	var initialized bool
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, ret *interface{}) {
		*ret = v0

		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					// skip first item(reflect.Type)
					if !initialized {
						initialized = true
						continue
					}

					v1 := reflect.ValueOf(*ret)
					v2 := reflect.ValueOf(v)
					*ret = funcValue.Call([]reflect.Value{v1, v2})[0].Interface()
				} else {
					(*wg).Done()
					return
				}
			default:
				continue
			}
		}
	}(&wg, &ret)
	wg.Wait()

	return &gollection{
		val: ret,
	}

}
