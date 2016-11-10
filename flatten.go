package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Flatten() *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.flattenStream()
	}
	return g.flatten()
}

func (g *gollection) flatten() *gollection {
	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Flatten called with non-slice value of type %T", g.slice),
		}
	}

	currentType := reflect.TypeOf(g.slice).Elem()
	if currentType.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Flatten called with non-slice-of-slice value of type %T", g.slice),
		}
	}

	// init
	ret := reflect.MakeSlice(currentType, 0, sv.Len())

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i).Interface()
		svv := reflect.ValueOf(v)
		for j := 0; j < svv.Len(); j++ {
			ret = reflect.Append(ret, svv.Index(j))
		}
	}

	return &gollection{
		slice: ret.Interface(),
		err:   nil,
	}
}

func (g *gollection) flattenStream() *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	go func() {
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					svv := reflect.ValueOf(v)
					for j := 0; j < svv.Len(); j++ {
						next.ch <- svv.Index(j).Interface()
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
