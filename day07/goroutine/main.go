package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10; i++ {
		//go hello(i) // 开启一个单独的goroutine去执行hello函数（任务）

		// 无名称的匿名函数（任务）
		// 不传参数时打印最多是10；传参打印无序
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main")
	// main 函数结束了，用main函数启动的goroutine也都结束了
	time.Sleep(time.Second)

}
