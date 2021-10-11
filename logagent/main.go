package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"study/logagent/kafka"
	"study/logagent/taillog"
	"time"
)

// logAgent入口程序
func run() {
	// 1、读取日志
	// 2、发送到kafka
	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 0、加载配置文件
	//cfg,err := ini.Load("config/config.ini")
	//if err != nil{
	//	fmt.Printf("config open failed,err:%v\n",err)
	//	return
	//}
	// 1、初始化kafka连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2、打开日志文件准备收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v\n", err)
		return
	}
	fmt.Println("init taillog success")
	run()
}
