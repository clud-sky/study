package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)
	b = make(chan int) // 不带缓冲的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10 // 卡住了（需要增加进城去取值）
	fmt.Println("10发送到通道b中了。。。")
	wg.Wait()
}

// 带缓冲的
func bufChannel() {
	fmt.Println(b)
	b = make(chan int, 2)
	b <- 10
	fmt.Println("10发送到通道b中了。。。")
	b <- 20
	fmt.Println("20发送到通道b中了。。。")
	x := <-b
	fmt.Println("从通道b中取到了", x)
	x = <-b
	fmt.Println("从通道b中取到了", x)
	close(b)
}

func main() {
	bufChannel()
}
