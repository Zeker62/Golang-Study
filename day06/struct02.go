package main

import "fmt"

type person01 struct {
	name string
	age int
	sex string
}

func main() {
	var p1 person01
	p1.name="张三"
	p1.age=23
	fmt.Println(p1.name,p1.sex,p1.age)
	fmt.Printf("p1=%#v",p1)
}
