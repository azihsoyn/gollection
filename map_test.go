package gollection_test

import (
	"fmt"
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	res, err := gollection.New(arr).Map(func(v int) int {
		return v * 2
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestMap_WithCast(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []string{"2", "4", "6", "8", "10", "12", "14", "16", "18", "20"}

	res, err := gollection.New(arr).Map(func(v int) string {
		return fmt.Sprintf("%d", v*2)
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestMap_EmptySlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New([]int{}).Map(func(v interface{}) interface{} {
		return ""
	}).Result()
	assert.NoError(err)
}
func TestMap_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Map(func(v interface{}) interface{} {
		return ""
	}).Result()
	assert.Error(err)
}
func TestMap_NotFunc(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New([]int{}).Map(0).Result()
	assert.Error(err)
}

func TestMap_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").
		Map(func(v interface{}) interface{} {
		return ""
	}).
		Map(func(v interface{}) interface{} {
		return ""
	}).
		Result()
	assert.Error(err)
}
