package main

import "fmt"

func main() {
	// 1、&：取地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)

	// *：根据地址取值
	m := *p
	fmt.Println(m)

	//var a *int // nil 不知道地址是什么所以错误
	//*a = 100
	//fmt.Println(*a)

	var b = new(int)
	*b = 100
	fmt.Println(*b)
	fmt.Println(&b)
}
