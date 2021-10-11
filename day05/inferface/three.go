package main

import "fmt"

// 接口的实现
type animal interface {
	move()
	eat(something string)
}

type cat1 struct {
	name string
	feet int8
}

func (c cat1) move() {
	fmt.Println("猫走猫步")
}

func (c cat1) eat(food string) {
	fmt.Printf("吃鱼%s...\n", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡冻")
}

func (c chicken) eat(food string) {
	fmt.Printf("吃虫子%s...\n", food)
}

func main() {
	var a1 animal
	bc := cat1{
		name: "淘气",
		feet: 4,
	}

	a1 = bc
	a1.eat("小黄鱼")
	fmt.Println(a1)

	kfc := chicken{feet: 10}

	a1 = kfc
	a1.eat("公鸡")
	fmt.Println(a1)
}
