package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Distinct() *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.distinctStream()
	}

	return g.distinct()

}

func (g *gollection) DistinctBy(f /*func(v <T1>) <T2>*/ interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.distinctByStream(f)
	}

	return g.distinctBy(f)
}

func (g *gollection) distinct() *gollection {
	sv, err := g.validateSlice("Distinct")
	if err != nil {
		return &gollection{err: err}
	}

	ret := reflect.MakeSlice(sv.Type(), 0, sv.Len())
	m := make(map[interface{}]bool)

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		id := v.Interface()
		if _, ok := m[id]; !ok {
			ret = reflect.Append(ret, v)
			m[id] = true
		}
	}

	return &gollection{
		slice: ret.Interface(),
	}
}

func (g *gollection) distinctStream() *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	var initialized bool
	m := make(map[interface{}]bool)
	go func() {
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					if !initialized {
						// initialize next stream type
						next.ch <- v.(reflect.Type)
						initialized = true
						continue
					}

					if _, ok := m[v]; !ok {
						next.ch <- v
						m[v] = true
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

func (g *gollection) validateDistinctByFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 {
		return reflect.Value{}, nil, fmt.Errorf("gollection.DistinctBy called with invalid func. required func(in <T>) out <T> but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) distinctBy(f interface{}) *gollection {
	sv, err := g.validateSlice("DistinctBy")
	if err != nil {
		return &gollection{err: err}
	}

	funcValue, funcType, err := g.validateDistinctByFunc(f)
	if err != nil {
		return &gollection{err: err}
	}

	resultSliceType := reflect.SliceOf(funcType.In(0))
	ret := reflect.MakeSlice(resultSliceType, 0, sv.Len())
	m := make(map[interface{}]bool)

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		id := funcValue.Call([]reflect.Value{v})[0].Interface()
		if _, ok := m[id]; !ok {
			ret = reflect.Append(ret, v)
			m[id] = true
		}
	}

	return &gollection{
		slice: ret.Interface(),
	}
}

func (g *gollection) distinctByStream(f interface{}) *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	funcValue, _, err := g.validateDistinctByFunc(f)
	if err != nil {
		return &gollection{err: err}
	}

	var initialized bool
	m := make(map[interface{}]bool)
	go func() {
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					if !initialized {
						// initialize next stream type
						next.ch <- v.(reflect.Type)
						initialized = true
						continue
					}

					id := funcValue.Call([]reflect.Value{reflect.ValueOf(v)})[0].Interface()
					if _, ok := m[id]; !ok {
						next.ch <- v
						m[id] = true
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
