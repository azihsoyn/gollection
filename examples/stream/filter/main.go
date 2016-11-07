package main

import (
	"fmt"
	"time"

	"github.com/azihsoyn/gollection"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("start  :", time.Now())
	res, _ := gollection.NewStream(arr).Filter(func(v int) bool {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("process 1st filter")
		return true
	}).Filter(func(v int) bool {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("process 2st filter")
		return true
	}).Result()
	fmt.Println("origin : ", arr)
	fmt.Println("ret    : ", res)
	fmt.Println("end    :", time.Now())
}
