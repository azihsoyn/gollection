package gollection_test

import (
	"testing"
	"time"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{1, 2, 3}

	res, err := gollection.New(arr).Take(3).Result()
	assert.NoError(err)
	assert.Equal(expect, res)

	expect = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res, err = gollection.New(arr).Take(30).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestTake_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Take(0).Result()
	assert.Error(err)
}

func TestTake_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").Take(0).Take(0).Result()
	assert.Error(err)
}

func TestTake_Stream(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{1, 2, 3}

	res, err := gollection.NewStream(arr).Take(3).Result()
	assert.NoError(err)
	assert.Equal(expect, res)

	expect = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res, err = gollection.NewStream(arr).Take(30).Result()
	assert.NoError(err)
	assert.Equal(expect, res)

	gollection.NewStream(arr).Filter(func(v int) bool {
		time.Sleep(1)
		return true
	}).Take(100)
}

func TestTake_Stream_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").Take(0).Result()
	assert.Error(err)
}

func TestTake_Stream_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.NewStream("not slice value").Take(0).Take(0).Result()
	assert.Error(err)
}
