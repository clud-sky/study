package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int
	lock   sync.Mutex
	wg     sync.WaitGroup
	rwLock sync.RWMutex
)

func read() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	rwLock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	fmt.Println(time.Now().Sub(start))
}
