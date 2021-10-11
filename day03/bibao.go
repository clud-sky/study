package main

import "fmt"

func main() {
	ret := lixiang(bibiaof2, 100, 200)
	ret()
}

func bibiaof2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func lixiang(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
	//x(m,n)
}
