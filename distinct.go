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

func (g *gollection) DistinctBy(f interface{}) *gollection {
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

	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.DistinctBy called with invalid func. required func(in <T>) out <T> but supplied %v", g.slice),
		}
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
