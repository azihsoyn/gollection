package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) FlatMap(f func(v interface{}) interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.FlatMap called with non-slice value of type %T", g.slice),
		}
	}

	ret := make([]interface{}, 0, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i).Interface()
		svv := reflect.ValueOf(v)
		if svv.Kind() != reflect.Slice {
			continue
		}
		for j := 0; j < svv.Len(); j++ {
			ret = append(ret, f(svv.Index(j).Interface()))
		}
	}

	return &gollection{
		slice: ret,
		err:   nil,
	}
}
