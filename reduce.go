package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Reduce(f interface{}) *gollection {
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
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Reduce called with empty slice of type %T", g.slice),
		}
	} else if sv.Len() == 1 {
		return &gollection{
			val: sv.Index(0).Interface(),
		}
	}

	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 2 || funcType.NumOut() != 1 {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Reduce called with invalid func. required func(in1, in2 <T>) out <T> but supplied %v", g.slice),
		}
	}

	ret := sv.Index(0).Interface()
	for i := 1; i < sv.Len(); i++ {
		v1 := reflect.ValueOf(ret)
		v2 := sv.Index(i)
		ret = funcValue.Call([]reflect.Value{v1, v2})[0].Interface()
	}

	return &gollection{
		val: ret,
	}
}
