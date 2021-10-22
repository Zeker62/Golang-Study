package main

import "fmt"

func show_data_10(a interface{}){
	fmt.Printf("通过interface接收了值为%v ，%T类型的变量\n",a,a)
}

func main() {
	var a interface{}
	b:=1000
	a=b
	show_data_10(a)
	c:=&b //指针类型
	a=c
	show_data_10(a)
	d:="字符串"
	a=d
	show_data_10(a)
}
