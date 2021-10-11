package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套
type animal4 interface {
	mover4
	eater4
}

type mover4 interface {
	move4()
}

type eater4 interface {
	eat4(string)
}

type cat4 struct {
	name string
	feet int8
}

func (c *cat4) mover() {
	fmt.Println("走猫步。。。。")
}

// cat4实现了eater接口
func (c *cat4) eat(food string) {
	fmt.Printf("猫可以吃%s....", food)
}

func main() {
	assign1(false)
}

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("类型错了")
	} else {
		fmt.Println(str)
	}
}

func assign1(a interface{}) {
	fmt.Printf("%T\n", a)

	switch t := a.(type) {
	case string:
		fmt.Print("是一个字符串：", t)
	case int:
		fmt.Print("是一个数字：", t)
	case bool:
		fmt.Print("是一个布尔", t)
	}
}
