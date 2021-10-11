package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

type Config struct {
	MysqlConfig `ini:"Mysql"`
	RedisConfig `ini:"Redis"`
}

func loadIni(filename string, data interface{}) error {
	// 0、参数校验
	// 0.1、传进来的data参数必须是指针类型（因为需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	//fmt.Println(t,t.Kind())
	if t.Kind() != reflect.Ptr {
		err := errors.New("data should be a pointer")
		return err
	}
	// 0.2、传进来的data参数是结构体类型指针（因为配置文件中各种键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err := errors.New("data param should be a struct pointer")
		return err
	}
	// 1、读文件得到字节类型
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// 将字节类型的文件内容转换为字符串
	lineSlice := strings.Split(string(b), "\r\n")
	//fmt.Printf("%#v\n", lineSlice)
	// 2、一行一行的读数据
	var structName string
	for idx, line := range lineSlice {
		// 2.1、如果是注释就跳过
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2、如果是[]开头的就表示是节
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err := errors.New(fmt.Sprintf("line:%d syntax error", idx+1))
				return err
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err := errors.New(fmt.Sprintf("line:%d syntax error", idx+1))
				return err
			}
			// 根据字符串sectionName需要结构题
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					//fmt.Printf("找到%s对应的嵌套结构体%s\n",sectionName,structName)
				}
			}
		} else {
			// 2.3、如果不是[开头就是=分割的键值对
			// 2.4、以等号分割这一行，等号左边是key,右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				return errors.New(fmt.Sprintf("line:%d syntax error", idx+1))
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 2.5、根据structName 去data里面把对应的嵌套结构体取出
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()
			if sType.Kind() != reflect.Struct {
				return errors.New(fmt.Sprintf("data中的%s字段应该是一个结构体", structName))
			}
			// 2.6、遍历嵌套结构体的每一个字符，判断tag是不是等于key
			var fieldName string
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i)
				if filed.Tag.Get("ini") == key {
					// 找到对应的字段
					fieldName = filed.Name
					break
				}
			}
			// 2.7、如果key=tag,赋值
			if len(fieldName) == 0 {
				continue
			}
			// 2.7.1 根据fieldName 去取出这个字段
			fileObj := sValue.FieldByName(fieldName)
			switch fileObj.Type().Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valueInt, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return errors.New(fmt.Sprintf("line:%d value type failed,err:%v\n", idx+1, err))
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				valueBool, err := strconv.ParseBool(value)
				if err != nil {
					return errors.New(fmt.Sprintf("line:%d value type failed,err:%v\n", idx+1, err))
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				valueFloat, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return errors.New(fmt.Sprintf("line:%d value type failed,err:%v\n", idx+1, err))
				}
				fileObj.SetFloat(valueFloat)
			}
			//fmt.Println(fileObj)
		}
	}

	return nil
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
	}
	fmt.Println(cfg.MysqlConfig.Address, cfg.RedisConfig.Test)
}
