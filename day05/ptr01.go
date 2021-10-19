package main

import "fmt"

func main() {
	a:=10
	b:=10
	c:=b
	fmt.Printf("变量a的值为：%v 地址为：%p\n",a,&a)
	fmt.Printf("变量b的值为: %v 地址为：%p \n",b,&b)
//	%p占位符是取地址占位符
	fmt.Printf("变量c的值为: %v 地址为：%p",c,&c)
}
