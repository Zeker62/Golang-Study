package main

import "fmt"

func main() {
	var array[3]int
	var array_string_2=[...]string{"上海","北京"}
	var array_bool[3]bool
	var array_int=[...]int{1,2}
	fmt.Println(array)
	fmt.Println("不确定初始化",array_string_2)
	fmt.Println(array_bool)
	fmt.Println("不确定初始化：",array_int)
}