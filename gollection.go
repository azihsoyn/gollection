package gollection

import (
	"fmt"
	"reflect"
	"sort"

	"go4.org/reflectutil"
)

type gollection struct {
	slice interface{}
	val   interface{}
	err   error
}

func New(slice interface{}) *gollection {
	return &gollection{
		slice: slice,
	}
}

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

func (g *gollection) Filter(f func(v interface{}) bool) *gollection {
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

	ret := make([]interface{}, 0, sv.Len())

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i).Interface()
		if f(v) {
			ret = append(ret, v)
		}
	}

	return &gollection{
		slice: ret,
	}
}

func (g *gollection) Map(f func(v interface{}) interface{}) *gollection {
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
	ret := make([]interface{}, 0, sv.Len())

	for i := 0; i < sv.Len(); i++ {
		v := f(sv.Index(i).Interface())
		ret = append(ret, v)
	}

	return &gollection{
		slice: ret,
	}
}

func (g *gollection) Reduce(f func(v1, v2 interface{}) interface{}) *gollection {
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

func (g *gollection) Result() (interface{}, error) {
	if g.val != nil {
		return g.val, g.err
	}
	return g.slice, g.err
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
