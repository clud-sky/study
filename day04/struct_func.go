package main

import "fmt"

type structPerson struct {
	name string
	age  int
}

// 构造函数：默认约定new开头
// 返回的结构体还是结构体指针
//当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *structPerson {
	return &structPerson{
		name: name,
		age:  age,
	}
}

// 表示符：变量名 函数名 类型名 方法名
// Go语言中如果标识符首字母是大写的，就表示对外部可见（暴露的，公有的）

// Dog 这是一个狗的结构体
type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{name: name}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~", d.name)
}

// 使用值接受者外部改不了（传递拷贝值进去）
func (p structPerson) guonian() {
	p.age++
}

// 使用指针接收者（传递内存地址进去）
func (p *structPerson) zhengguonian() {
	p.age++
}

func main() {
	p1 := newPerson("元帅", 18)
	//p2 := newPerson("之恩",88)
	//fmt.Println(p1,p2)
	//fmt.Println(p1.name)
	//d1 := newDog("周六")
	//fmt.Println(d1)
	//
	//d2 := newDog("旺财")
	//d2.wang()

	//p1.guonian()
	//fmt.Println(p1.age)
	p1.zhengguonian()
	fmt.Println(p1.age)
}
