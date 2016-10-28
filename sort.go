package gollection

import (
	"fmt"
	"reflect"
	"sort"

	"go4.org/reflectutil"
)

func (g *gollection) SortBy(f interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.SortBy called with non-slice value of type %T", g.slice),
		}
	}
	orig := reflect.MakeSlice(sv.Type(), sv.Len(), sv.Cap())
	ret := reflect.MakeSlice(sv.Type(), sv.Len(), sv.Cap())
	reflect.Copy(orig, sv)

	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 2 || funcType.NumOut() != 1 || funcType.Out(0).Kind() != reflect.Bool {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.SortBy called with invalid func. required func(in1, in2 <T>) bool but supplied %v", g.slice),
		}
	}

	less := func(i, j int) bool {
		return funcValue.Call([]reflect.Value{sv.Index(i), sv.Index(j)})[0].Interface().(bool)
	}

	sort.Sort(&funcs{
		length: sv.Len(),
		less:   less,
		swap:   reflectutil.Swapper(g.slice),
	})
	reflect.Copy(ret, sv)
	reflect.Copy(sv, orig)

	return &gollection{
		slice: ret.Interface(),
		err:   nil,
	}
}

type funcs struct {
	length int
	less   func(i, j int) bool
	swap   func(i, j int)
}

func (f *funcs) Len() int           { return f.length }
func (f *funcs) Less(i, j int) bool { return f.less(i, j) }
func (f *funcs) Swap(i, j int) {
	f.swap(i, j)
}
