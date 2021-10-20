package main

import "fmt"

type person05 struct{
	name string
	age int
	sex string
}
func main() {
	p2:=&person05{}
	p2.sex="nv"
	p2.name="Emil"
	fmt.Printf("type of p2 : %T\n",p2)
	fmt.Printf("value of p2: %v\n",p2)
	fmt.Printf("value of p2: %v\n",*p2)
	(*p2).age=22
	(*p2).name="Have"
	fmt.Printf("value of p2: %v\n",*p2)
}
