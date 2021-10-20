package main

import "fmt"

type person10 struct {
	name string
	age  int
}

func newPerson10(name string,age int) *person10{
	return &person10{
		name:name,
		age:age,
	}
}
func main() {
	p10:=newPerson10("张三",23)
	fmt.Printf("value of p10: %#v",*p10)
}
