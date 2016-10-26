package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Flatten() *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Flatten called with non-slice value of type %T", g.slice),
		}
	}

	var ret reflect.Value
	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i).Interface()
		svv := reflect.ValueOf(v)
		// init
		if i == 0 {
			ret = reflect.MakeSlice(svv.Type(), 0, sv.Len())
		}
		if svv.Kind() != reflect.Slice {
			continue
		}
		for j := 0; j < svv.Len(); j++ {
			ret = reflect.Append(ret, svv.Index(j))
		}
	}

	return &gollection{
		slice: ret.Interface(),
		err:   nil,
	}
}
