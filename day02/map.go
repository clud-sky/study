package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	testMapSlice()
}

func testMapSlice() {
	// 元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "A"
	fmt.Println(s1)

	// 值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["成都"] = []int{10, 20, 30}
	fmt.Println(m1)

}

func testMap2() {
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 0)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	// 取出map中的所有key存入切片keys
	var keys = make([]int, 0)
	for _, v := range scoreMap {
		keys = append(keys, v)
	}
	sort.Ints(keys)
	fmt.Println(keys)
}

func testMap1() {
	var m1 map[string]int
	fmt.Println(m1 == nil) // 还没有初始化（没有在内存中开辟空间）

	m1 = make(map[string]int, 5) // 要估算好该map容量，避免运行期间扩容
	m1["理想"] = 10
	m1["sjkiiii"] = 35
	m1["a"] = 36
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	// 判断
	value, ok := m1["jslj"]
	if !ok {
		fmt.Println(value)
		fmt.Println("没有该key")
	} else {
		fmt.Println(value)
	}

	// map遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	fmt.Println("----")
	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
		fmt.Println(m1[k])
	}
	fmt.Println("----")
	// 只遍历value
	// 删除
	delete(m1, "a")
	for _, v := range m1 {
		fmt.Println(v)
	}
}
