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
