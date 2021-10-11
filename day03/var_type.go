package main

import "fmt"

var x = 100

func main() {
	//fmt.Printf("%T\n",testF1)
	//fmt.Printf("%T\n",testType)
	fmt.Println(testType2(testType))
}

// 函数还可以昨晚返回值
func testType2(x func() int) func(int, int) int {
	return ff
}

func ff(a, b int) int {
	return a + b
}

// 函数也可以作为参数的类型
func testType1(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func testType() int {
	return 10
}

// 变量先在函数内部寻找在到函数外找
func testF1() {
	fmt.Println(x)
}
