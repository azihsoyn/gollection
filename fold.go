package gollection

import (
	"reflect"
	"sync"
)

func (g *gollection) Fold(v0 interface{}, f /* func(v1, v2 <T>) <T> */ interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.foldStream(v0, f)
	}

	return g.fold(v0, f)
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
		ret = processReduceFunc(funcValue, v1, v2).Interface()
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
					*ret = processReduceFunc(funcValue, v1, v2).Interface()
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
