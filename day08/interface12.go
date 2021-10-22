package main

import "fmt"

func main() {
	var a interface{}
	a="aabbcc"
	v,ok:=a.(string)
	if ok{
		fmt.Println(v)
	}else{
		fmt.Println("断言错误")
	}
}
