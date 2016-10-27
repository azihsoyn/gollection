package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Take(n int) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Take called with non-slice value of type %T", g.slice),
		}
	}

	limit := sv.Len()
	if n < limit {
		limit = n
	}

	ret := reflect.MakeSlice(sv.Type(), 0, sv.Len())

	for i := 0; i < limit; i++ {
		ret = reflect.Append(ret, sv.Index(i))
	}

	return &gollection{
		slice: ret.Interface(),
	}
}
