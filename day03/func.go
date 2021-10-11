package main

import "fmt"

func main() {
	fmt.Println(f5())
}

func f5() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// 5
func f4() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 返回的是X
func f3() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// 5
func f2() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f1() {
	defer fmt.Println("欢迎您")
	defer fmt.Println("人民")
	fmt.Println("hello 成都！")
}
