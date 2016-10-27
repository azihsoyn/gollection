package gollection_test

import (
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	assert := assert.New(t)

	arr := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9, 10},
	}
	expect := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	res, err := gollection.New(arr).FlatMap(func(v interface{}) interface{} {
		if n, ok := v.(int); ok {
			return n * 2
		}
		return ""
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestFlatMap_InterfaceSlice(t *testing.T) {
	assert := assert.New(t)
	arr := []interface{}{
		[]int{1, 2, 3},
		"a", "b",
		nil,
	}
	expect := []int{2, 4, 6}

	res, err := gollection.New(arr).FlatMap(func(v interface{}) interface{} {
		if n, ok := v.(int); ok {
			return n * 2
		}
		return ""
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestFlatMap_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").FlatMap(func(v interface{}) interface{} {
		return ""
	}).Result()
	assert.Error(err)
}

func TestFlatMap_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").
		FlatMap(func(v interface{}) interface{} {
		return ""
	}).
		FlatMap(func(v interface{}) interface{} {
		return ""
	}).
		Result()
	assert.Error(err)
}
