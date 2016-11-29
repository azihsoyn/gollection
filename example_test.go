package gollection_test

import (
	"fmt"

	"github.com/azihsoyn/gollection"
)

func Example_distinct() {
	arr := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10}

	res, err := gollection.New(arr).Distinct().Result()
	fmt.Println(res, err)
	// Output: [1 2 3 4 5 6 7 8 9 10] <nil>
}

func Example_filter() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res, err := gollection.New(arr).Filter(func(v int) bool {
		return v > 5
	}).Result()
	fmt.Println(res, err)
	// Output: [6 7 8 9 10] <nil>
}

func Example_flatMap() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9, 10},
	}

	res, err := gollection.New(arr).FlatMap(func(v int) int {
		return v * 2
	}).Result()
	fmt.Println(res, err)
	// Output: [2 4 6 8 10 12 14 16 18 20] <nil>
}

func Example_flatten() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9, 10},
	}

	res, err := gollection.New(arr).Flatten().Result()
	fmt.Println(res, err)
	// Output: [1 2 3 4 5 6 7 8 9 10] <nil>
}

func Example_fold() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res, err := gollection.New(arr).Fold(100, func(v1, v2 int) int {
		return v1 + v2
	}).Result()
	fmt.Println(res, err)
	// Output: 155 <nil>
}

func Example_map() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res, err := gollection.New(arr).Map(func(v int) int {
		return v * 2
	}).Result()
	fmt.Println(res, err)
	// Output: [2 4 6 8 10 12 14 16 18 20] <nil>
}

func Example_reduce() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res, err := gollection.New(arr).Reduce(func(v1, v2 int) int {
		return v1 + v2
	}).Result()
	fmt.Println(res, err)
	// Output: 55 <nil>
}

func Example_sort() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	res, err := gollection.New(arr).SortBy(func(v1, v2 int) bool {
		return v1 < v2
	}).Result()
	fmt.Println(arr)
	fmt.Println(res, err)
	// Output: [1 2 3 4 5 6 7 8 9 10 9 8 7 6 5 4 3 2 1]
	// [1 1 2 2 3 3 4 4 5 5 6 6 7 7 8 8 9 9 10] <nil>
}

func Example_take() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res, err := gollection.New(arr).Take(3).Result()
	fmt.Println(res, err)
	res, err = gollection.New(arr).Take(30).Result()
	fmt.Println(res, err)
	// Output: [1 2 3] <nil>
	// [1 2 3 4 5 6 7 8 9 10] <nil>
}

func Example_customType() {
	type User struct {
		ID   int
		Name string
	}

	in := []User{
		{ID: 1, Name: "aaa"},
		{ID: 2, Name: "bbb"},
		{ID: 3, Name: "ccc"},
		{ID: 4, Name: "ddd"},
		{ID: 5, Name: "eee"},
		{ID: 6, Name: "fff"},
		{ID: 7, Name: "ggg"},
	}
	res, err := gollection.New(in).Filter(func(v User) bool {
		return v.ID > 5
	}).Result()
	fmt.Printf("%#v %#v\n", res.([]User), err)
	// Output: []gollection_test.User{gollection_test.User{ID:6, Name:"fff"}, gollection_test.User{ID:7, Name:"ggg"}} <nil>
}

func Example_customTypeWithResultAs() {
	type User struct {
		ID   int
		Name string
	}

	in := []User{
		{ID: 1, Name: "aaa"},
		{ID: 2, Name: "bbb"},
		{ID: 3, Name: "ccc"},
		{ID: 4, Name: "ddd"},
		{ID: 5, Name: "eee"},
		{ID: 6, Name: "fff"},
		{ID: 7, Name: "ggg"},
	}
	var res []User
	err := gollection.New(in).Filter(func(v User) bool {
		return v.ID > 5
	}).ResultAs(&res)
	fmt.Printf("%#v %#v\n", res, err)
	// Output: []gollection_test.User{gollection_test.User{ID:6, Name:"fff"}, gollection_test.User{ID:7, Name:"ggg"}} <nil>
}
