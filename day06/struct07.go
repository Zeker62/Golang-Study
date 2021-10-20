package main

import "fmt"

type person07 struct{
	name string
	age int
	sex string
}
func main() {
	//对结构体初始化
	p1:=person07{
		name:"Emak",
		age:22,
		sex:"male",
	}
	fmt.Printf("value of p1:%v\n",p1)
	//对结构体指针初始化
	p2:=&person07{
		name:"Alice",
		sex:"female",
	}
	fmt.Printf("value of p2:%v\n",*p2)
	//省略键的结构体初始化赋值
	p3:=&person07{
		"Pbied",
		22,
		"male",
	}
	fmt.Printf("value of p3:%v",*p3)
}
