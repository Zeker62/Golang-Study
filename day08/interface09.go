package main

import "fmt"

func main() {
	var inter interface{}
	s:="这是一个字符串"
	inter=s
	fmt.Println(inter)
	a:=10
	inter=a
	fmt.Println(inter)
	c:=100.0002
	inter=c
	fmt.Println(inter)
}

//这是一个字符串
//10
//100.0002