package gollection_test

import (
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	assert := assert.New(t)
	arr := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9, 10},
	}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res, err := gollection.New(arr).Flatten().Result()
	assert.NoError(err)
	assert.Equal(expect, res)

	arr = [][]int{
		{1},
		nil,
		{2},
	}
	expect = []int{1, 2}

	res, err = gollection.New(arr).Flatten().Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestFlatten_EmptySlice(t *testing.T) {
	assert := assert.New(t)
	arr := [][]int{}
	expect := []int{}

	res, err := gollection.New(arr).Flatten().Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestFlatten_InterfaceSlice(t *testing.T) {
	assert := assert.New(t)
	arr := [][]interface{}{
		[]interface{}{1, 2, 3},
		[]interface{}{"a", "b"},
		nil,
	}
	expect := []interface{}{1, 2, 3, "a", "b"}

	res, err := gollection.New(arr).Flatten().Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestFlatten_NotSlice(t *testing.T) {
	assert := assert.New(t)

	_, err := gollection.New("not slice value").Flatten().Result()
	assert.Error(err)
}

func TestFlatten_HavingError(t *testing.T) {
	assert := assert.New(t)

	_, err := gollection.New("not slice value").Flatten().Flatten().Result()
	assert.Error(err)
}
