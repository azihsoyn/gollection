package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Reduce(f func(v1, v2 interface{}) interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Reduce called with non-slice value of type %T", g.slice),
		}
	}

	if sv.Len() == 0 {
		return &gollection{}
	} else if sv.Len() == 1 {
		return &gollection{
			val: sv.Index(0).Interface(),
		}
	}

	v1 := sv.Index(0).Interface()
	for i := 1; i < sv.Len(); i++ {
		v2 := sv.Index(i).Interface()
		v1 = f(v1, v2)
	}

	return &gollection{
		val: v1,
	}
}
