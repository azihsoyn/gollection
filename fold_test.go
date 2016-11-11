package gollection_test

import (
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := 155

	res, err := gollection.New(arr).Fold(100, func(v1, v2 int) int {
		return v1 + v2
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)

	arr = []int{}
	expect = 100

	res, err = gollection.New(arr).Fold(100, func(v1, v2 int) int {
		return v1 + v2
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}
func TestFold_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Fold(100, func(v1, v2 interface{}) interface{} {
		return ""
	}).Result()
	assert.Error(err)
}
func TestFold_NotFunc(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New([]int{0}).Fold(100, 0).Result()
	assert.Error(err)
}

func TestFold_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").
		Fold(100, func(v1, v2 interface{}) interface{} {
		return ""
	}).
		Fold(100, func(v1, v2 interface{}) interface{} {
		return ""
	}).
		Result()
	assert.Error(err)
}

func TestFold_Stream(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := 155

	res, err := gollection.NewStream(arr).Fold(100, func(v1, v2 int) int {
		return v1 + v2
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)

	arr = []int{}
	expect = 100

	res, err = gollection.NewStream(arr).Fold(100, func(v1, v2 int) int {
		return v1 + v2
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}
func TestFold_Stream_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").Fold(100, func(v1, v2 interface{}) interface{} {
		return ""
	}).Result()
	assert.Error(err)
}
func TestFold_Stream_NotFunc(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream([]int{0}).Fold(100, 0).Result()
	assert.Error(err)
}

func TestFold_Stream_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").
		Fold(100, func(v1, v2 interface{}) interface{} {
		return ""
	}).
		Fold(100, func(v1, v2 interface{}) interface{} {
		return ""
	}).
		Result()
	assert.Error(err)
}
