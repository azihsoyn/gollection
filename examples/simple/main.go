package main

import (
	"fmt"

	"github.com/azihsoyn/gollection"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	res, _ := gollection.New(arr).Sort(func(i, j int) bool {
		return arr[i] < arr[j]
	}).Filter(func(v interface{}) bool {
		if n, ok := v.(int); ok && n > 5 {
			return true
		}
		return false
	}).Result()
	fmt.Println("origin : ", arr)
	fmt.Println("ret : ", res)
}
