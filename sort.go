package gollection

import (
	"fmt"
	"reflect"
	"sort"

	"go4.org/reflectutil"
)

func (g *gollection) Sort(less func(i, j int) bool) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return &gollection{
			slice: nil,
			err:   fmt.Errorf("gollection.Sort called with non-slice value of type %T", g.slice),
		}
	}
	orig := reflect.MakeSlice(sv.Type(), sv.Len(), sv.Cap())
	ret := reflect.MakeSlice(sv.Type(), sv.Len(), sv.Cap())
	reflect.Copy(orig, sv)

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
