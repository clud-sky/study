package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

type Cat struct {
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)

	var c = Cat{}
	reflectType(c)

	// 获取值类型
	vv := reflect.ValueOf(&b)
	va := vv.Kind()
	fmt.Println(va)

	vv.Elem().SetInt(200)
	fmt.Println(b)
}
