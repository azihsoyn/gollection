package gollection_test

import (
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	res, err := gollection.New(arr).Map(func(v interface{}) interface{} {
		if n, ok := v.(int); ok {
			return n * 2
		}
		return ""
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}
