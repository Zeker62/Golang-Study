package main

import "fmt"

type person06 struct{
	name string
	age int
	sex string
}
func main() {
	var p1 person06
	fmt.Printf("value of p1: %#v",p1)
}
