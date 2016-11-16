package gollection

import (
	"reflect"
	"sort"

	"go4.org/reflectutil"
)

func processDistinct(v interface{}, m map[interface{}]bool) bool {
	if _, ok := m[v]; !ok {
		m[v] = true
		return true
	}
	return false
}

func processDistinctBy(fv, v reflect.Value, m map[interface{}]bool) bool {
	id := fv.Call([]reflect.Value{v})[0].Interface()
	if _, ok := m[id]; !ok {
		m[id] = true
		return true
	}
	return false
}

func processFilter(fv, v reflect.Value) bool {
	return fv.Call([]reflect.Value{v})[0].Interface().(bool)
}

func processMapFunc(fv, arg reflect.Value) reflect.Value {
	return fv.Call([]reflect.Value{arg})[0]
}

func processReduceFunc(fv, arg1, arg2 reflect.Value) reflect.Value {
	return fv.Call([]reflect.Value{arg1, arg2})[0]
}

func processSort(fv, ret reflect.Value) {
	less := func(i, j int) bool {
		return fv.Call([]reflect.Value{ret.Index(i), ret.Index(j)})[0].Interface().(bool)
	}

	sort.Sort(&funcs{
		length: ret.Len(),
		less:   less,
		swap:   reflectutil.Swapper(ret.Interface()),
	})
}
