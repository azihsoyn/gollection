package gollection

import (
	"fmt"
	"reflect"
	"sort"
	"sync"

	"go4.org/reflectutil"
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

func (g *gollection) validateSortByFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 2 ||
		funcType.NumOut() != 1 ||
		funcType.Out(0).Kind() != reflect.Bool {
		return reflect.Value{}, nil, fmt.Errorf("gollection.SortBy called with invalid func. required func(in1, in2 <T>) bool but supplied %v", funcType)
	}
	return funcValue, funcType, nil
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

	less := func(i, j int) bool {
		return funcValue.Call([]reflect.Value{ret.Index(i), ret.Index(j)})[0].Interface().(bool)
	}

	sort.Sort(&funcs{
		length: sv.Len(),
		less:   less,
		swap:   reflectutil.Swapper(ret.Interface()),
	})

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

	less := func(i, j int) bool {
		return funcValue.Call([]reflect.Value{ret.Index(i), ret.Index(j)})[0].Interface().(bool)
	}

	sv := reflect.ValueOf(ret.Interface())

	sort.Sort(&funcs{
		length: sv.Len(),
		less:   less,
		swap:   reflectutil.Swapper(ret.Interface()),
	})

	go func() {
		// initialze next stream type
		next.ch <- currentType

		for i := 0; i < sv.Len(); i++ {
			next.ch <- sv.Index(i).Interface()
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
