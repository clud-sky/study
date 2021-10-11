package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine
	for i := 0; i < 3; i++ {
		go worker(i+1, jobs, results)
	}

	// 5个任务向goroutine 传值
	for i := 0; i < 5; i++ {
		jobs <- i + 1
	}

	close(jobs)

	// 输出结果
	for j := 0; j < 5; j++ {
		<-results
	}
}
