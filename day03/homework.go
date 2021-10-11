package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	testHuiw()
}

// 回文判断
func testHuiw() {
	var s = "上海2自来水来自海上"
	r := make([]rune, 0, len(s))
	for _, c := range s {
		r = append(r, c)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("No")
			return
		}
	}
}

// 判断相同单词次数
func testCi() {
	s1 := "how do you do"
	str := strings.Split(s1, " ")
	fmt.Println(str)
	m1 := make(map[string]int, 10)
	for _, v := range str {
		fmt.Println(v)
		if _, ok := m1[v]; !ok {
			m1[v] = 1
		} else {
			m1[v]++
		}

	}
	fmt.Println(m1)

	for k, v := range m1 {
		if v > 1 {
			fmt.Println(k)
		}
	}
}

// 判断字符串中的汉字数量
func testHan() {
	s1 := "hello沙河"
	count := 0
	for _, v := range s1 {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println(count)
}
