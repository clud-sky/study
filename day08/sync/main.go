package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

var (
	x    int
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	for i := 0; i < 5; i++ {
		lock.Lock()
		x = x + 1
		fmt.Println(x)
		lock.Unlock()
	}
	defer wg.Done()
}

func main() {
	value := strconv.FormatFloat(9.285, 'f', 2, 64)
	fmt.Println(value)
	return
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	time.Sleep(time.Second * 10)
	fmt.Println(x)

	sss := 3.1415
	fmt.Println(math.Round(sss))
}

func round(x float64) float64 {
	return float64(math.Floor(x + 0/5))
}
