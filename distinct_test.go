package gollection_test

import (
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 1, 2, 3}
	expect := []int{1, 2, 3}

	res, err := gollection.New(arr).Distinct().Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestDistinct_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Distinct().Result()
	assert.Error(err)
}

func TestDistinct_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Distinct().Distinct().Result()
	assert.Error(err)
}
