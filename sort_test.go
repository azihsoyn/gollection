package gollection_test

import (
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expect := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10}
	original := make([]int, len(arr))
	copy(original, arr)

	res, err := gollection.New(arr).Sort(func(i, j int) bool {
		return arr[i] < arr[j]
	}).Result()

	assert.NoError(err)
	assert.Equal(expect, res)
	// check not changed
	assert.Equal(original, arr)
}

func TestSort_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Sort(func(i, j int) bool {
		return false
	}).Result()
	assert.Error(err)
}

func TestSort_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").
		Sort(func(i, j int) bool {
		return false
	}).
		Sort(func(i, j int) bool {
		return false
	}).
		Result()
	assert.Error(err)
}
