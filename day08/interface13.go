package main

import "fmt"

func switchtype(a interface{}){
	switch v:=a.(type){
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {
	var a interface{}
	a="aavvss"
	switchtype(a)
}
//x is a string，value is aavvss