package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed,err:%v\n", err)
		return
	}

	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2019-08-04 14:12:12", loc)
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}

	fmt.Println(timeObj)
}
