package main

import (
	"fmt"
	"os"
)

// 学生信息结构体
type student struct {
	id int
	name string
	age int
	sex string
	grade *grade //使用指针传递，可以减少空间开销

}
// 学生成绩结构体
type grade struct {
	Chinese int
	English int
	Math int
}

// 构造函数：
func NewStudent(id int,name string,sex string,age int,Chinese int,English int,Math int) *student{
	s:=new(student)
	s.id=id
	s.name=name
	s.sex=sex
	s.age=age
	s.grade=&grade{
		Chinese:Chinese,
		English: English,
		Math:Math,
	}
	return s
}

var students=make(map[string]student,100) //定义一个容量为100，类型为学生的map

func Menu(){
	fmt.Println("+++++++++++++++++go语言学生管理系统+++++++++++++++")
	fmt.Println("+++++++++++++++++1.添加学生信息++++++++++++++++++")
	fmt.Println("+++++++++++++++++2.查看学生信息++++++++++++++++++")
	fmt.Println("+++++++++++++++++3.修改学生信息++++++++++++++++++")
	fmt.Println("+++++++++++++++++4.学生信息公示++++++++++++++++++")
	fmt.Println("+++++++++++++++++0.退出学生系统++++++++++++++++++")
}

func AddStudent(){
	var id int
	var name string
	var age int
	var sex string
	var Chinese int
	var English int
	var Math int
	fmt.Println("-----------请输入学生信息----------")
	fmt.Println("输入学生学号")
	fmt.Scanf("%d\n",&id)
	fmt.Println("请输入姓名： ")
	fmt.Scanf("%s\n",&name)
	fmt.Println("请输出年龄： ")
	fmt.Scanf("%d\n",&age)
	fmt.Println("请输入性别： ")
	fmt.Scanf("%s\n",&sex)
	fmt.Println("请分别输入语文、英语、数学成绩，并用空格隔开：")
	fmt.Scanf("%d %d %d",&Chinese,&English,&Math)
	newStudent:=NewStudent(id,name,sex,age,Chinese,English,Math)
	fmt.Printf("当前输入学生信息:%v",newStudent)

	students[newStudent.name]=*newStudent
	fmt.Println("添加成功")

}

func SearchStudent(name string) bool{

	value,ok:=students[name]
	if ok{
		fmt.Println("该学生存在")
		fmt.Println("该学生的信息如下:")
		fmt.Printf("学号: %d\n姓名: %s\n性别: %s\n年龄: %d\n--成绩--\n语文成绩:%d\n数学成绩:%d\n英语成绩:%d\n",
			value.id,value.name,value.sex,value.age,value.grade.Chinese,value.grade.Math,value.grade.English)
		return true
	}else {
		fmt.Println("该学生不存在")
		return false
	}

}
func ModifyStudent(name string) bool {
	_,ok:=students[name]
	if ok {
		var id int
		var name string
		var age int
		var sex string
		var Chinese int
		var English int
		var Math int
		fmt.Println("该学生存在,请输入修改学生的参数")
		fmt.Println("-----------请输入学生信息----------")
		fmt.Println("输入学生学号")
		fmt.Scanf("%d\n",&id)
		fmt.Println("请输入姓名： ")
		fmt.Scanf("%s\n",&name)
		fmt.Println("请输出年龄： ")
		fmt.Scanf("%d\n",&age)
		fmt.Println("请输入性别： ")
		fmt.Scanf("%s\n",&sex)
		fmt.Println("请分别输入语文、英语、数学成绩，并用空格隔开：")
		fmt.Scanf("%d %d %d",&Chinese,&English,&Math)
		newStudent:=NewStudent(id,name,sex,age,Chinese,English,Math)
		fmt.Printf("当前输入学生信息:%v",newStudent)
		students[name]=*newStudent
		fmt.Println("添加成功")
		return true
	} else {
		fmt.Println("该学生不存在")
		return false
	}
}


//输出所有学生的成绩
func showStudent(){

	for _,value:=range students{
		fmt.Printf("学号: %d \t姓名: %s\t性别: %s\t年龄: %d\t--成绩--\t语文成绩:%d\t数学成绩:%d\t英语成绩:%d\t",
			value.id,value.name,value.sex,value.age,value.grade.Chinese,value.grade.Math,value.grade.English)
		fmt.Println()
	}

}
func main() {
	for{
		Menu()
		var name string
		var option int
		fmt.Printf("\n请输入你要执行的操作")
		fmt.Scanf("%d\n",&option)
		switch  option {
		case 1:
			AddStudent()
		case 2:
			fmt.Print("请输入你要查询的学生姓名")
			fmt.Scanf("%s",&name)
			SearchStudent(name)
		case 3:
			fmt.Print("请输入你要修改的学生姓名:")
			fmt.Scanf("%s", &name)
			ModifyStudent(name)
		case 4:
			showStudent()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("输入有误,请重新输入")
		}
	}
}
