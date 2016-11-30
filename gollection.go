/*
Package gollection provides collection util to any type slices.
*/
package gollection

import (
	"fmt"
	"reflect"
	"sync"
)

type interfaceSlice struct {
	slice []interface{}
}

func (s *interfaceSlice) truncate() {
	s.slice = s.slice[:0]
}

func (s *interfaceSlice) free() {
	slicePool.Put(s)
}

func newInterfaceSlice() *interfaceSlice {
	s := slicePool.Get().(*interfaceSlice)
	s.truncate()
	return s
}

var slicePool = sync.Pool{
	New: func() interface{} {
		return &interfaceSlice{
			slice: make([]interface{}, 0, 10),
		}
	},
}

type gollection struct {
	slice  interface{}
	val    interface{}
	ch     chan interface{}
	err    error
	slice2 *interfaceSlice
	meta   meta
}

type meta struct {
	Len  int
	Type reflect.Type
}

// New returns a gollection instance which can method chain *sequentially* specified by some type of slice.
func New(slice interface{}) *gollection {
	return &gollection{
		slice: slice,
	}
}

func New2(slice interface{}) *gollection {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{err: fmt.Errorf("gollection.%s called with non-slice value of type %T", slice)}
	}
	s := newInterfaceSlice()

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i).Interface()
		s.slice = append(s.slice, v)
	}
	return &gollection{
		slice:  slice,
		slice2: s,
		meta: meta{
			Len:  sv.Len(),
			Type: sv.Type(),
		},
	}
}

// Result return a collection processed value and error.
func (g *gollection) Result() (interface{}, error) {
	if g.ch != nil {
		return g.resultStream()
	}
	if g.slice2 != nil {
		ret := reflect.MakeSlice(g.meta.Type, 0, g.meta.Len)
		for _, v := range g.slice2.slice {
			ret = reflect.Append(ret, reflect.ValueOf(v))
		}
		return ret.Interface(), g.err
	}
	return g.result()
}

func (g *gollection) result() (interface{}, error) {
	if g.val != nil {
		return g.val, g.err
	}
	return g.slice, g.err
}

func (g *gollection) ResultAs(out interface{}) error {
	if g.err != nil {
		return g.err
	}

	iv := reflect.ValueOf(g.slice)
	if g.val != nil {
		iv = reflect.ValueOf(g.val)
	}

	ov := reflect.ValueOf(out)
	if ov.Kind() != reflect.Ptr || iv.Type() != ov.Elem().Type() {
		return fmt.Errorf("gollection.ResultAs called with unexpected type %T, expected %s", g.slice, ov.Elem().Type())
	}

	ov.Elem().Set(iv)
	return nil
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
