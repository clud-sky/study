package main

import (
	"fmt"
	"sort"
)

func main() {
	testLx()
}

func testLx() {
	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:])
	fmt.Println(a1)
}

// copy()
func testCopy() {
	a1 := []int{1, 2, 3, 5}
	a2 := a1 // 赋值
	x1 := a1[:]
	var a3 = make([]int, 4)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)

	// 将a1中的索引为1的元素删掉
	a1 = append(a1[:1], a1[2:]...)
	fmt.Printf("a1=%v len(a1)=%d cap(a1)=%d\n", a1, len(a1), cap(a1))

	fmt.Println(x1)
}

// append()为切片追加元素
//追加元素后，原来的底层数组放不下的时候，go底层就会把底层数组换一个；
//容量在小于1024时直接翻倍；如果大于1024时每次增加25%
//如果容量溢出，会使用最开始申请的容量；扩容会根据元素类型做不同处理，方式也不一样
//必须用变量接收
func testAppend() {
	s1 := []string{"金牛", "青羊", "高新（西）"}
	s1 = append(s1, "成华", "温江")
	fmt.Println(s1)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	ss := []string{"双流", "金堂", "新都"}
	s1 = append(s1, ss...)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}

// 切片遍历
func testSlice() {
	s1 := []int{1, 2, 3, 4}
	for i, v := range s1 {
		fmt.Println(i, v)
	}
}

// make()函数创造切片
func testMake() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2:%v len(s2):%d cap(s2):%d\n", s2, len(s2), cap(s2))

	s3 := make([]int, 0)
	fmt.Printf("s3:%v len(s3):%d cap(s3):%d\n", s3, len(s3), cap(s3))
}

// 切片
func testCap() {
	s := "大的覅偶很为dldjofiwe"
	s1 := s[:3]
	fmt.Println(len(s1), s1)

	var s2 []int
	var s3 []string
	fmt.Println(s2, s3)

	s2 = []int{1, 2, 3}
	s3 = []string{"金牛", "青羊", "高新（西）"}
	fmt.Println(s2, s3)
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))
	fmt.Printf("len(s3):%d cap(s3):%d\n", len(s3), cap(s3))

	s4 := s3[1:]
	s3[1] = "成华" // 切片是引用，与更改一样针对的是底层数字;数据都是保存在底层数组里面
	fmt.Println(s4)
	s4[1] = "脾毒"
	fmt.Println(s4)
}

// 运算符
func testYsf() {
	var (
		a = 5
		b = 2
	)
	// 算数运算
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// 关系运算
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

}
