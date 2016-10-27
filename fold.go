package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Fold(v0 interface{}, f func(v1, v2 interface{}) interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Fold called with non-slice value of type %T", g.slice),
		}
	}

	if sv.Len() < 1 {
		return &gollection{
			val: v0,
		}
	}

	v1 := v0
	for i := 0; i < sv.Len(); i++ {
		v2 := sv.Index(i).Interface()
		v1 = f(v1, v2)
	}

	return &gollection{
		val: v1,
	}
}
