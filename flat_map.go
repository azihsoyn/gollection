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

	// init
	retType := reflect.ValueOf(f(nil)).Type()
	ret := reflect.MakeSlice(reflect.SliceOf(retType), 0, sv.Len())

	// avoid "panic: reflect: call of reflect.Value.Interface on zero Value"
	// see https://github.com/azihsoyn/gollection/issues/7
	if sv.Len() == 0 {
		return &gollection{
			slice: ret.Interface(),
		}
	}

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i).Interface()
		svv := reflect.ValueOf(v)
		if svv.Kind() != reflect.Slice {
			continue
		}
		for j := 0; j < svv.Len(); j++ {
			v := reflect.ValueOf(f(svv.Index(j).Interface()))
			ret = reflect.Append(ret, v)
		}
	}

	return &gollection{
		slice: ret.Interface(),
		err:   nil,
	}
}
