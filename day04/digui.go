package main

import "fmt"

func main() {
	fmt.Println(taijie(4))
}

// 计算N的阶层
func diguiF1(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * diguiF1(n-1)
}

// n个台阶，一次可以走一步也可以走两步，有多少中方法
func taijie(n int64) int64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return taijie(n-1) + taijie(n-2)
}
