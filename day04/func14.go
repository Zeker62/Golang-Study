package main

import (
	"fmt"
	"strings"
)

// 这是一个添加后缀名的例子
func add13(s string)func (string)string{
	return func(s2 string)string{
		if !strings.HasSuffix(s2,s){
		//	该函数判断的是s是否是在s2的尾部
		//	s是后缀名，s2是文件名
			return s2+s
		}else{
			return s2
		}
	}
}
func main() {
	var f=add13(".jpg")
	fmt.Println(f("123"))
	fmt.Println(f("456.jpg"))
	var z=add13(".txt")
	fmt.Println(z("hhh.txt"))
}
