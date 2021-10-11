package main

import (
	"testlog/mylogger"
)

func main() {
	//log := mylogger.NewLog("Info")
	log := mylogger.NewFileLogger("Info", "./", "zhouling.log", 100*1024)

	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
	}
}
