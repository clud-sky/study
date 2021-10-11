package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	testswitch()
}

// 循环
func testswitch() {
	n := 2
	switch n {
	case 2:
		fmt.Println(2)
		goto ss // 调到指定标签
	case 7:
		fmt.Println(7)
		fallthrough // 向下运行一个case
	default:
		fmt.Println(0)
	}

ss: // 标签
	fmt.Println(465461)
}

// 乘法口诀
func testSf() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

// 修改字符串
func testRune() {
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[1] = '大'
	fmt.Println(string(s3))
}

// 遍历成字符
func testRane() {
	s := "沙河快递呢hello"
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
}

func testString() {
	//ch := make(chan int,1)
	//for i:=0;i<10;i++{
	//	select {
	//	case x := <-ch:
	//		fmt.Println(x)
	//	case ch<- i:
	//	}
	//}

	f1 := float32(1.23456)
	fmt.Printf("%T\n", f1)
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%2f\n", math.Pi)

	f2 := "12345978642"
	ass := strings.Contains(f2, "0")
	fmt.Println(ass)
}
