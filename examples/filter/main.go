package main

import (
	"fmt"

	"github.com/azihsoyn/gollection"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res, _ := gollection.New(arr).Filter(func(v interface{}) bool {
		if n, ok := v.(int); ok && n > 5 {
			return true
		}
		return false
	}).Result()
	fmt.Println("origin : ", arr)
	fmt.Println("ret    : ", res)
}
