package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) validateSlice(funcName string) (reflect.Value, error) {
	sv := reflect.ValueOf(g.slice)
	if sv.Kind() != reflect.Slice {
		return reflect.Value{}, fmt.Errorf("gollection.%s called with non-slice value of type %T", funcName, g.slice)
	}
	return sv, nil
}

func (g *gollection) validateSliceOfSlice(funcName string) (reflect.Type, error) {
	currentType := reflect.TypeOf(g.slice).Elem()
	if currentType.Kind() != reflect.Slice {
		return nil, fmt.Errorf("gollection.%s called with non-slice value of type %T", funcName, g.slice)
	}
	return currentType, nil
}

func (g *gollection) validateDistinctByFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 {
		return reflect.Value{}, nil, fmt.Errorf("gollection.DistinctBy called with invalid func. required func(in <T>) out <T> but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) validateFilterFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 1 ||
		funcType.NumOut() != 1 ||
		funcType.Out(0).Kind() != reflect.Bool {
		return reflect.Value{}, nil, fmt.Errorf("gollection.Filter called with invalid func. required func(in <T>) bool but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) validateFlatMapFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 {
		return reflect.Value{}, nil, fmt.Errorf("gollection.FlatMap called with invalid func. required func(in <T>) out <T> but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) validateFoldFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 2 ||
		funcType.NumOut() != 1 {
		return reflect.Value{}, nil, fmt.Errorf("gollection.Fold called with invalid func. required func(in1, in2 <T>) out <T> but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) validateMapFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 1 ||
		funcType.NumOut() != 1 {
		return reflect.Value{}, nil, fmt.Errorf("gollection.Map called with invalid func. required func(in <T>) out <T> but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) validateReduceFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 2 ||
		funcType.NumOut() != 1 {
		return reflect.Value{}, nil, fmt.Errorf("gollection.Reduce called with invalid func. required func(in1, in2 <T>) out <T> but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (g *gollection) validateSortByFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 2 ||
		funcType.NumOut() != 1 ||
		funcType.Out(0).Kind() != reflect.Bool {
		return reflect.Value{}, nil, fmt.Errorf("gollection.SortBy called with invalid func. required func(in1, in2 <T>) bool but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}
