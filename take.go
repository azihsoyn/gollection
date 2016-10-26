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
			err:   fmt.Errorf("gollection.Filter called with non-slice value of type %T", g.slice),
		}
	}

	limit := sv.Len()
	if n < limit {
		limit = n
	}

	ret := make([]interface{}, 0, limit)

	for i := 0; i < limit; i++ {
		v := sv.Index(i).Interface()
		ret = append(ret, v)
	}

	return &gollection{
		slice: ret,
	}
}
