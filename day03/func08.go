package main

import "fmt"

type calculate func(int,int)int
//定义这种形态的函数都是calculate类型

func sum0001(a int,b int)int{
	s:=a+b
	return s
}
func sub00001(x int,y int)int{
	s:=x-y
	return s
}


func main() {
	var c calculate
//	定义一个变量，类型是calculate样式的函数
	c=sum0001
//	可以将函数sum赋值给c
	fmt.Printf("type of c is %T\n",c)
	fmt.Println(c(10,20))

	// 没有经过calculate直接赋值
	f:=sub00001
	fmt.Printf("type of f is %T\n",f)
	fmt.Println(f(20,10))
}
