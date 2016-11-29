package main

import (
	"fmt"

	"github.com/azihsoyn/gollection"
)

type User struct {
	ID   int
	Name string
}

func main() {
	in := []User{
		{ID: 1, Name: "aaa"},
		{ID: 2, Name: "bbb"},
		{ID: 3, Name: "ccc"},
		{ID: 4, Name: "ddd"},
		{ID: 5, Name: "eee"},
		{ID: 6, Name: "fff"},
		{ID: 7, Name: "ggg"},
	}
	var out []User
	err := gollection.New(in).Filter(func(v User) bool {
		return v.ID > 5
	}).ResultAs(&out)
	fmt.Printf("out with ResultAs       : %#v\n", out)
	fmt.Println("err : ", err)

	ret, _ := gollection.New(in).Filter(func(v User) bool {
		return v.ID > 5
	}).Result()
	out = ret.([]User)
	fmt.Printf("out with type assertion : %#v\n", out)
}
