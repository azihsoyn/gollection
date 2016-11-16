package gollection

import (
	"fmt"
	"reflect"
	"sync"
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

func (g *gollection) validateReduceFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 2 ||
		funcType.NumOut() != 1 {
		return reflect.Value{}, nil, fmt.Errorf("gollection.Reduce called with invalid func. required func(in1, in2 <T>) out <T> but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) reduce(f interface{}) *gollection {
	sv, err := g.validateSlice("Reduce")
	if err != nil {
		return &gollection{err: err}
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

	funcValue, _, err := g.validateReduceFunc(f)
	if err != nil {
		return &gollection{err: err}
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
	funcValue, _, err := g.validateReduceFunc(f)
	if err != nil {
		return &gollection{err: err}
	}

	var ret interface{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, ret *interface{}, err *error) {
		var initialized bool
		var skippedFirst bool
		var itemNum int
		var currentType reflect.Type

		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					// skip first item(reflect.Type)
					if !skippedFirst {
						skippedFirst = true
						currentType = v.(reflect.Type)
						continue
					}

					if !initialized {
						itemNum++
						*ret = reflect.ValueOf(v).Interface()
						initialized = true
						continue
					}

					v1 := reflect.ValueOf(*ret)
					v2 := reflect.ValueOf(v)
					*ret = funcValue.Call([]reflect.Value{v1, v2})[0].Interface()
				} else {
					if itemNum == 0 {
						*err = fmt.Errorf("gollection.Reduce called with empty slice of type %s", currentType)
					}
					(*wg).Done()
					return
				}
			default:
				continue
			}
		}
	}(&wg, &ret, &err)
	wg.Wait()

	if err != nil {
		return &gollection{
			err: err,
		}
	}

	return &gollection{
		val: ret,
	}
}
