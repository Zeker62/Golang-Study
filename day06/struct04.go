package main

import "fmt"

type person04 struct{
	name string
	age int
	sex string
}
func main() {
	var p1=new(person04)
	p1.sex="女"
	p1.name="小红"
	p1.age=20
	fmt.Printf("p1: %#v\n",p1)
	fmt.Printf("value of p1:%v\n",*p1)
	fmt.Printf("value of p1:%p\n",p1)
	fmt.Printf("address of p1:%p\n",&p1)
	fmt.Printf("address of *p1:%p\n",&*p1)

	var p2 person04
	p2.name="王五"
	p2.sex="男"
	fmt.Printf("value of p2:%v\n",p2)
	fmt.Printf("address of p2:%p\n",&p2)
	fmt.Printf("value of p2.name:%v\n",p2.name)
}
