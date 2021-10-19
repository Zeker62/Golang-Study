package main

import "fmt"

func ptr06_a(x int){
	x=20
	fmt.Println("在函数ptr06_a中x的地址",&x)
}
func ptr06_b(x *int){
	// x此时是一个指针
	*x=20
	fmt.Println("在函数ptr06_b中x的地址",&x)
	fmt.Println("在函数ptr06_b中*x的地址",&*x)
}
func main() {
	x:=100
	fmt.Println("全局变量x的地址：",&x)
	ptr06_a(x)
	fmt.Println(x)
	ptr06_b(&x) //*int是指针类型，指针的值是地址，所以需要传递地址
	fmt.Println(x)
	fmt.Println("最后x的地址：",&x)
}
