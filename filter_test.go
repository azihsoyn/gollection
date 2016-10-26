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

	res, err := gollection.New(arr).Filter(func(v interface{}) bool {
		if n, ok := v.(int); ok && n > 5 {
			return true
		}
		return false
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}
