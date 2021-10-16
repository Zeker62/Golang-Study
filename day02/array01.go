package main

import "fmt"

func main() {
	// 数组定义
	var array[3]int
	var array_string[3]string
	var array_string_2=[4]string{"上海","北京"}
	var array_bool[3]bool
	var array_int=[3]int{1,2}
	fmt.Println(array)
	fmt.Println(array_string)
	fmt.Println("部分初始化",array_string_2)
	fmt.Println(array_bool)
	fmt.Println("部分初始化：",array_int)
}
//[0 0 0]
//[  ]
//部分初始化 [上海 北京  ]
//[false false false]
//部分初始化： [1 2 0]