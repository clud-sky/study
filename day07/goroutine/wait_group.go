package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    // int64
		r2 := rand.Intn(10) // 0<=x<10
		fmt.Println(r1, r2)
	}
}

//
func f1(i int) {
	defer wg.Done() // 完成计数器-1
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	//f()

	for i := 0; i < 10; i++ {
		wg.Add(1) // 计数器每次+1
		go f1(i)
	}

	// 如何知道这10个goroutine都结束了
	wg.Wait() // 等待wg的计数器减为0
}
