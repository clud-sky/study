package main

import "fmt"

// 结构体
type person struct {
	Name   string
	Age    int
	Gender string
	Hobby  []string
}

func main() {
	var zhoulin person
	zhoulin.Name = "周林"
	zhoulin.Age = 9
	zhoulin.Gender = "男"
	zhoulin.Hobby = []string{"篮球", "皮球"}

	fmt.Println(zhoulin)

	var s struct {
		name string
		age  int
	}

	s.name = "的后代"
	s.age = 1
	fmt.Print(s)
}
