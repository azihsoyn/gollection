package main

import (
	"fmt"

	"github.com/azihsoyn/gollection"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	res, _ := gollection.New(arr).SortBy(func(v1, v2 interface{}) bool {
		return v1.(int) < v2.(int)
	}).Result()
	fmt.Println("origin : ", arr)
	fmt.Println("ret    : ", res)
}
