package main

import "fmt"

func main() {
	var array=[...]int{1,2,3,4,5,6,7,8,9,10}
	//len函数可以获取数组长度
	fmt.Println("使用for循环遍历数组array")
	for i:=0;i<len(array);i++{
		fmt.Print(array[i]," ")
	}
	fmt.Println("\n使用for range循环遍历数组")
	for _,a:=range array{
		fmt.Print(a," ")
	}
}
//使用for循环遍历数组array
//1 2 3 4 5 6 7 8 9 10
//使用for range循环遍历数组
//1 2 3 4 5 6 7 8 9 10