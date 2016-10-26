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
}
