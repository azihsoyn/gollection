package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Filter(f func(v interface{}) bool) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Filter called with non-slice value of type %T", g.slice),
		}
	}

	ret := reflect.MakeSlice(sv.Type(), 0, sv.Len())

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if f(v.Interface()) {
			ret = reflect.Append(ret, sv.Index(i))
		}
	}

	return &gollection{
		slice: ret.Interface(),
	}
}