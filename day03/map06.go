package main

import "fmt"

func main() {
	// 切片里面有map
	slice :=make([]map[string]string,3)
	for a,b :=range slice {
		fmt.Printf("key:%d value:%v\n",a,b)
	}
	fmt.Println("初始化……")
	slice[0] =make(map[string]string,4)
	slice[0]["name"]="aaa"
	slice[0]["age"]="123"
	for index,value := range slice {
		fmt.Printf("key:%d value:%v\n",index,value)
	}
	fmt.Println(slice)
}
