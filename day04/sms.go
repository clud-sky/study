package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看\新增\删除
*/

type student struct {
	id   int64
	name string
}

var (
	allStudent map[int64]*student // 变量声明
)

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	// 把所有学生打印出来
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

func addStudent() {
	// 向allstutdent中添加一个新的学生
	// 1、创建一个新学生
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)

	// 2、造学生(调用student的构造函数)
	newStu := newStudent(id, name)

	// 追加到哦allStudent这个map中
	allStudent[id] = newStu

}

func deleteStudent() {
	// 请输入要删除的学号
	var (
		deleteID int64
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&deleteID)
	// 去allStudent这个map中根据学号删除对应的键值对
	delete(allStudent, deleteID)
}

func main() {
	allStudent = make(map[int64]*student, 48) // 初始化（开辟内存空间）
	for {
		// 打印菜单
		fmt.Println("欢迎光临学生管理系统")
		fmt.Println(`
		1、查看所有学生
		2、新增学生
		3、删除学生
		4、退出
	`)
		fmt.Print("请输入你要干啥：")

		// 等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项！\n", choice)

		// 执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("去，乱搞~")

		}
	}
}
