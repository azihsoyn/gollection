/*
Package gollection provides collection util to any type slices.
*/
package gollection

import (
	"fmt"
	"reflect"
	"sync"
)

type gollection struct {
	slice interface{}
	val   interface{}
	ch    chan interface{}
	err   error
}

// New returns a gollection instance which can method chain *sequentially* specified by some type of slice.
func New(slice interface{}) *gollection {
	return &gollection{
		slice: slice,
	}
}

func (g *gollection) validateSlice(funcName string) (reflect.Value, error) {
	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return reflect.Value{}, fmt.Errorf("gollection.%s called with non-slice value of type %T", funcName, g.slice)
	}
	return sv, nil
}

func (g *gollection) validateSliceOfSlice(funcName string) error {
	currentType := reflect.TypeOf(g.slice).Elem()
	if currentType.Kind() != reflect.Slice {
		return fmt.Errorf("gollection.%s called with non-slice value of type %T", funcName, g.slice)
	}
	return nil
}

// Result return a collection processed value and error.
func (g *gollection) Result() (interface{}, error) {
	if g.ch != nil {
		return g.resultStream()
	}
	return g.result()
}

func (g *gollection) result() (interface{}, error) {
	if g.val != nil {
		return g.val, g.err
	}
	return g.slice, g.err
}

func (g *gollection) resultStream() (interface{}, error) {
	var ret reflect.Value
	var initialized bool
	var err error

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(err *error) {
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					if e, ok := v.(error); ok {
						*err = e
						wg.Done()
						return
					}
					if !initialized {
						ret = reflect.MakeSlice(v.(reflect.Type), 0, 0)
						initialized = true
						continue
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
	}(&err)
	wg.Wait()

	if err != nil {
		return nil, err
	}

	return ret.Interface(), nil
}

// NewStream returns a gollection instance which can method chain *parallel* specified by some type of slice.
func NewStream(slice interface{}) *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			err: fmt.Errorf("gollection.NewStream called with non-slice value of type %T", slice),
		}
	}

	go func() {
		// initialze next stream type
		next.ch <- sv.Type()

		for i := 0; i < sv.Len(); i++ {
			next.ch <- sv.Index(i).Interface()
		}
		close(next.ch)
	}()
	return next
}
