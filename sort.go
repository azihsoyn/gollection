package gollection

import (
	"reflect"
	"sync"
)

func (g *gollection) SortBy(f /* func(v1, v2 <T>) bool */ interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.sortByStream(f)
	}

	return g.sortBy(f)
}

func (g *gollection) sortBy(f interface{}) *gollection {
	sv, err := g.validateSlice("SortBy")
	if err != nil {
		return &gollection{err: err}
	}

	ret := reflect.MakeSlice(sv.Type(), sv.Len(), sv.Cap())
	reflect.Copy(ret, sv)

	funcValue, _, err := g.validateSortByFunc(f)
	if err != nil {
		return &gollection{err: err}
	}

	processSort(funcValue, ret)

	return &gollection{
		slice: ret.Interface(),
		err:   nil,
	}
}

func (g *gollection) sortByStream(f interface{}) *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	funcValue, _, err := g.validateSortByFunc(f)
	if err != nil {
		return &gollection{err: err}
	}

	var ret reflect.Value
	var initialized bool
	var skippedFirst bool
	var currentType reflect.Type

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, currentType *reflect.Type) {
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					if !skippedFirst {
						skippedFirst = true
						*currentType = v.(reflect.Type)
						continue
					}

					// initialze next stream type
					if !initialized {
						ret = reflect.MakeSlice(reflect.SliceOf(reflect.ValueOf(v).Type()), 0, 0)
						initialized = true
					}

					ret = reflect.Append(ret, reflect.ValueOf(v))
				} else {
					wg.Done()
					return
				}
			default:
				continue
			}
		}
	}(&wg, &currentType)
	wg.Wait()

	processSort(funcValue, ret)

	go func() {
		// initialze next stream type
		next.ch <- currentType

		for i := 0; i < ret.Len(); i++ {
			next.ch <- ret.Index(i).Interface()
		}
		close(next.ch)
	}()
	return next
}

type funcs struct {
	length int
	less   func(i, j int) bool
	swap   func(i, j int)
}

func (f *funcs) Len() int           { return f.length }
func (f *funcs) Less(i, j int) bool { return f.less(i, j) }
func (f *funcs) Swap(i, j int) {
	f.swap(i, j)
}
