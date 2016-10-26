package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Map(f func(v interface{}) interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Map called with non-slice value of type %T", g.slice),
		}
	}
	var ret reflect.Value

	for i := 0; i < sv.Len(); i++ {
		v := reflect.ValueOf(f(sv.Index(i).Interface()))
		// init
		if i == 0 {
			ret = reflect.MakeSlice(reflect.SliceOf(v.Type()), 0, sv.Len())
		}
		ret = reflect.Append(ret, v)
	}

	return &gollection{
		slice: ret.Interface(),
	}
}
