package main

import (
	"fmt"

	"github.com/azihsoyn/gollection"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res, _ := gollection.New(arr).Fold(100, func(v1, v2 interface{}) interface{} {
		n1, ok1 := v1.(int)
		n2, ok2 := v2.(int)
		if ok1 && ok2 {
			return n1 + n2
		}
		return ""
	}).Result()
	fmt.Println("origin : ", arr)
	fmt.Println("ret    : ", res)
}
