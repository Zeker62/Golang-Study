package main

import "fmt"

//类型定义和类型别名的区别
func main() {
	type MyInt int
	type cname_int=int

	var a MyInt
	var b cname_int
	fmt.Printf("Type of MyInt: %T\n",a)
	fmt.Printf("Type of cname_int: %T",b)

}
