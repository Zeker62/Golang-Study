package main

import "fmt"

func main() {
	a:=10
	b:=&a //b成了指针类型
	fmt.Printf("变量a的值为：%v 地址为：%p 类型为：%T\n",a,&a,a)
	fmt.Printf("变量b的值为: %v 地址为：%p 类型为：%T",b,&b,b)
	//	%p占位符是取地址占位符
}
