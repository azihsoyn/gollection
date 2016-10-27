package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Distinct() *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Distinct called with non-slice value of type %T", g.slice),
		}
	}

	ret := reflect.MakeSlice(sv.Type(), 0, sv.Len())
	m := make(map[interface{}]bool)

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if _, ok := m[v.Interface()]; !ok {
			ret = reflect.Append(ret, v)
			m[v.Interface()] = true
		}
	}

	return &gollection{
		slice: ret.Interface(),
	}
}

func (g *gollection) DistinctBy(f func(v interface{}) interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.DistinctBy called with non-slice value of type %T", g.slice),
		}
	}

	ret := reflect.MakeSlice(sv.Type(), 0, sv.Len())
	m := make(map[interface{}]bool)

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		id := f(v.Interface())
		if _, ok := m[id]; !ok {
			ret = reflect.Append(ret, v)
			m[id] = true
		}
	}

	return &gollection{
		slice: ret.Interface(),
	}
}
