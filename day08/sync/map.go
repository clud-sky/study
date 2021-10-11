package main

import (
	"fmt"
	"strconv"
	"sync"
)

// go内置的map不是并发安全的

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 31; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := strconv.Itoa(i)
			set(key, n)
			fmt.Printf("k=:%v,v=%v\n", key, n)
		}(i)
		wg.Wait()
	}
}

//var m2 = sync.Map{}
//
//func main(){
//	var wg sync.WaitGroup
//	for i := 0;i <25;i++{
//		wg.Add(1)
//		go func(n int) {
//			defer wg.Done()
//			key := strconv.Itoa(n)
//			m2.Store(key,n) // 必须使用sync.map内置的store方法设置键值对
//			value,_ := m2.Load(key) // 必须使用sync.map提供的load方法根据key取值
//			fmt.Printf("k=:%v,v=%v\n",key,value)
//		}(i)
//	}
//	wg.Wait()
//}
