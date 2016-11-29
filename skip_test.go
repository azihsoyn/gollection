package gollection_test

import (
	"testing"
	"time"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestSkip(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{4, 5, 6, 7, 8, 9, 10}

	res, err := gollection.New(arr).Skip(3).Result()
	assert.NoError(err)
	assert.Equal(expect, res)

	expect = []int{}
	res, err = gollection.New(arr).Skip(30).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestSkip_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Skip(0).Result()
	assert.Error(err)
}

func TestSkip_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Skip(0).Skip(0).Result()
	assert.Error(err)
}

func TestSkip_Stream(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{4, 5, 6, 7, 8, 9, 10}

	res, err := gollection.NewStream(arr).Skip(3).Result()
	assert.NoError(err)
	assert.Equal(expect, res)

	expect = []int{}
	res, err = gollection.NewStream(arr).Skip(30).Result()
	assert.NoError(err)
	assert.Equal(expect, res)

	gollection.NewStream(arr).Filter(func(v int) bool {
		time.Sleep(1)
		return true
	}).Skip(100)
}

func TestSkip_Stream_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").Skip(0).Result()
	assert.Error(err)
}

func TestSkip_Stream_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").Skip(0).Skip(0).Result()
	assert.Error(err)
}
