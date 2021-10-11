package mylogger

import (
	"fmt"
	"time"
)

// 往终端上写日志相关内容

// 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// 构造函数
func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (l ConsoleLogger) enable(logLevel LogLevel) bool {
	return l.Level <= logLevel
}

func (l ConsoleLogger) log(lv LogLevel, msg string, a ...interface{}) {
	if l.enable(lv) {
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(2)
		msg = fmt.Sprintf(msg, a...)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, getLogstring(lv), funcName, fileName, lineNo, msg)
	}
}

func (l ConsoleLogger) Debug(msg string) {
	if l.enable(DEBUG) {
		l.log(DEBUG, msg)
	}
}

func (l ConsoleLogger) Info(msg string) {
	if !l.enable(INFO) {
		return
	}
	l.log(INFO, msg)
}

func (l ConsoleLogger) Warning(msg string) {
	if !l.enable(WARNING) {
		return
	}
	l.log(WARNING, msg)
}

func (l ConsoleLogger) Error(msg string) {
	if !l.enable(ERROR) {
		return
	}
	l.log(ERROR, msg)
}

func (l ConsoleLogger) Fatal(msg string) {
	if !l.enable(FATAL) {
		return
	}
	l.log(FATAL, msg)
}
