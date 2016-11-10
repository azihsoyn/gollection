package gollection_test

import (
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{6, 7, 8, 9, 10}

	res, err := gollection.New(arr).Filter(func(v int) bool {
		return v > 5
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}
func TestFilter_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Filter(func(v interface{}) bool {
		return true
	}).Result()
	assert.Error(err)
}
func TestFilter_NotFunc(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New([]int{0}).Filter(0).Result()
	assert.Error(err)
}

func TestFilter_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").
		Filter(func(v interface{}) bool {
			return true
		}).
		Filter(func(v interface{}) bool {
			return true
		}).
		Result()
	assert.Error(err)
}

func TestFilter_Stream(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{6, 7, 8, 9, 10}

	res, err := gollection.NewStream(arr).Filter(func(v int) bool {
		return v > 5
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestFilter_Stream_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").Filter(func(v interface{}) bool {
		return true
	}).Result()
	assert.Error(err)
}

func TestFilter_Stream_NotFunc(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream([]int{0}).Filter(0).Result()
	assert.Error(err)
}

func TestFilter_Stream_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").
		Filter(func(v interface{}) bool {
			return true
		}).
		Filter(func(v interface{}) bool {
			return true
		}).
		Result()
	assert.Error(err)
}
