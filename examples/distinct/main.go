package main

import (
	"fmt"

	"github.com/azihsoyn/gollection"
)

func main() {
	arr := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10}

	res, _ := gollection.New(arr).Distinct().Result()
	fmt.Println("origin : ", arr)
	fmt.Println("ret    : ", res) // {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}
