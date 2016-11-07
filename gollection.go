package gollection

import (
	"fmt"
	"reflect"
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

func (g *gollection) Result() (interface{}, error) {
	if g.val != nil {
		return g.val, g.err
	}
	return g.slice, g.err
}

func (g *gollection) ResultAs(out interface{}) error {
	if g.val != nil {
		return g.err
	}
	iv := reflect.ValueOf(g.slice)
	ov := reflect.ValueOf(out)
	if ov.Kind() != reflect.Ptr || iv.Type() != ov.Elem().Type() {
		return fmt.Errorf("gollection.ResultAs called with unexpected type %T, expected %s", g.slice, ov.Elem().Type())
	}

	ov.Elem().Set(iv)
	return nil
}
