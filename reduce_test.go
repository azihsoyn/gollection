package gollection_test

import (
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	assert := assert.New(t)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := 55

	res, err := gollection.New(arr).Reduce(func(v1, v2 interface{}) interface{} {
		n1, ok1 := v1.(int)
		n2, ok2 := v2.(int)
		if ok1 && ok2 {
			return n1 + n2
		}
		return ""
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)

	arr = []int{1}
	expect = 1

	res, err = gollection.New(arr).Reduce(func(v1, v2 interface{}) interface{} {
		n1, ok1 := v1.(int)
		n2, ok2 := v2.(int)
		if ok1 && ok2 {
			return n1 + n2
		}
		return ""
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestReduce_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Reduce(func(v1, v2 interface{}) interface{} {
		return ""
	}).Result()
	assert.Error(err)
}

func TestReduce_EmptySlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New([]int{}).Reduce(func(v1, v2 interface{}) interface{} {
		return ""
	}).Result()
	assert.Error(err)
}

func TestReduce_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").
		Reduce(func(v1, v2 interface{}) interface{} {
		return ""
	}).
		Reduce(func(v1, v2 interface{}) interface{} {
		return ""
	}).
		Result()
	assert.Error(err)
}