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
	sv, err := g.validateSlice("Flatten")
	if err != nil {
		return &gollection{err: err}
	}

	currentType, err := g.validateSliceOfSlice("Flatten")
	if err != nil {
		return &gollection{err: err}
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

	var initialized bool
	go func() {
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					// initialze next stream type
					if !initialized {
						currentType := v.(reflect.Type).Elem()
						if currentType.Kind() != reflect.Slice {
							next.ch <- fmt.Errorf("gollection.Flatten called with non-slice-of-slice value of type %s", currentType)
						}
						next.ch <- currentType
						initialized = true
						continue
					}

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
