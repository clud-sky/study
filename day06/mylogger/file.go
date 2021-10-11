package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里写日志

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *LogMsg
}

type LogMsg struct {
	Level     LogLevel
	Msg       string
	FuncName  string
	FileName  string
	Timestamp string
	Line      int
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *LogMsg, 50000),
	}
	err = fl.initFile() //按照文件路劲喝文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

func (l *FileLogger) initFile() error {
	fullFileName := path.Join(l.filePath, l.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}

	l.fileObj = fileObj
	l.errFileObj = errFileObj
	// 开启一个后台的goroutine去写日志
	for i := 0; i < 5; i++ {
		go l.writeLogBackground()
	}
	return nil
}

func (l *FileLogger) close() {
	l.fileObj.Close()
	l.errFileObj.Close()
}

func (l *FileLogger) enable(logLevel LogLevel) bool {
	return l.Level <= logLevel
}

func (l *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= l.maxFileSize
}

func (l *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	logName := path.Join(l.filePath, l.fileName)
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	// 1、关闭当前的日志文件
	file.Close()
	// 2、备份一下 rename
	os.Rename(logName, newLogName)
	// 3、打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err
	}
	// 4、将打开的文建对象赋值给f.fileObj
	l.fileObj = fileObj
	return fileObj, nil
}

func (l *FileLogger) writeLogBackground() {
	for {
		select {
		case logTmp := <-l.logChan:
			fmt.Fprintf(l.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", logTmp.Timestamp, getLogstring(logTmp.Level), logTmp.FuncName, logTmp.FileName, logTmp.Line, logTmp.Msg)
			if logTmp.Level >= ERROR {
				// 如果要记录的日志大于等于ERROR级别，我还要再err日志文件中再记录一遍
				fmt.Fprintf(l.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", logTmp.Timestamp, getLogstring(logTmp.Level), logTmp.FuncName, logTmp.FileName, logTmp.Line, logTmp.Msg)
			}
		default:
			// 如果没有取到值就休息一会儿
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (l *FileLogger) log(lv LogLevel, msg string) {
	if l.enable(lv) {
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		// 先把日志发送到通道中
		// 1、造一个Msg对象
		logTmp := &LogMsg{
			Level:     lv,
			Msg:       msg,
			FuncName:  funcName,
			FileName:  fileName,
			Timestamp: now,
			Line:      lineNo,
		}
		select {
		case l.logChan <- logTmp:
		default:
			// 如果通道满了就把日志丢掉保证不出现柱塞
			time.Sleep(time.Millisecond * 500)
		}
		if l.checkSize(l.fileObj) {
			newFile, err := l.splitFile(l.fileObj)
			if err != nil {
				return
			}
			l.fileObj = newFile
		}
	}
}

func (l *FileLogger) Debug(msg string) {
	l.log(DEBUG, msg)
}

func (l *FileLogger) Info(msg string) {
	l.log(INFO, msg)
}

func (l *FileLogger) Warning(msg string) {
	l.log(WARNING, msg)
}

func (l *FileLogger) Error(msg string) {
	l.log(ERROR, msg)
}

func (l *FileLogger) Fatal(msg string) {
	l.log(FATAL, msg)
}
