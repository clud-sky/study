package main

import "fmt"

type persons struct {
	name, gender string
}

// go语言中函数参数永远传的是拷贝
func sf(x persons) {
	x.gender = "女" // 修改的是副本的gender
}

func sf1(x *persons) {
	//(*x).gender = "女" // 根据内存地址找到原变量，修改的原来的变量
	x.gender = "女"
}

func main() {
	var p persons
	p.name = "周林"
	p.gender = "男"
	sf(p)
	fmt.Println(p)

	sf1(&p)
	fmt.Print(p)

	var p2 = new(persons)
	p2.name = "2jhdf"
	fmt.Println(p2)
	fmt.Printf("%T\n", p2)
}
