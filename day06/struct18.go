package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID int
	Gender string
	Name string
}

type Class struct {
	Title string
	Students []*Student //定义结构体切片，因为一个班上会有很多学生，使用指针可以减少内存消耗
}


func main() {
	c:=&Class{
		Title: "计算机科学与技术1班",
		Students: make([]*Student,0,200), //结构体切片也需要初始化
	}
	for i:=0 ;i<10 ;i++{
		stu:=&Student{
			Name: fmt.Sprintf("student:%02d",i),
			Gender:"man",
			ID:i,
		}
		c.Students=append(c.Students,stu)
	}
	//JSON 序列化：即生成JSON格式字符串
	data,error:=json.Marshal(c)
	if error!=nil{
		fmt.Println("json序列化出现错误")
		return
	}
	fmt.Printf("jion:%s\n",data)

	//JSON反序列化
	str:=data
	c1:=&Class{} //创建一个空的Class对象，用于接收值
	error=json.Unmarshal([]byte(str),c1)
	if error!=nil{
		fmt.Println("json反序列化出现错误")
		return
	}
	fmt.Printf("%#v\n",c1)


}
