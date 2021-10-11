package main

import (
	"fmt"
	"study/day09/split_string"
)

func main() {
	str := split_string.Split("abdebesbser", "b")
	fmt.Printf("%#v\n", str)
	str1 := split_string.Split("bbbb", "b")
	fmt.Printf("%#v\n", str1)
}
