package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Filter(f interface{}) *gollection {
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

	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 || funcType.Out(0).Kind() != reflect.Bool {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Filter called with invalid func. required func(in <T>) bool but supplied %v", g.slice),
		}
	}

	resultSliceType := reflect.SliceOf(funcType.In(0))
	ret := reflect.MakeSlice(resultSliceType, 0, sv.Len())

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if funcValue.Call([]reflect.Value{v})[0].Interface().(bool) {
			ret = reflect.Append(ret, sv.Index(i))
		}
	}

	return &gollection{
		slice: ret.Interface(),
	}
}
