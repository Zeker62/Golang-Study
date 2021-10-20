package main

import "fmt"

func main() {
	var p1 struct{name string;age int}
	p1.name="李四"
	p1.age=33
	fmt.Printf("%#v\n",p1)
	fmt.Println(p1.name,p1.age)
}
