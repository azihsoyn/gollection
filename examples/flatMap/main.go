package main

import (
	"fmt"

	"github.com/azihsoyn/gollection"
)

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9, 10},
	}

	res, _ := gollection.New(arr).FlatMap(func(v interface{}) interface{} {
		if n, ok := v.(int); ok {
			return n * 2
		}
		return 0
	}).Result()
	fmt.Println("origin : ", arr)
	fmt.Println("ret    : ", res)
}
