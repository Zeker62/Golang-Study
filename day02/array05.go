package main

import "fmt"

func main() {
	//二维数组
	var array=[3][2]int{{1,2},{3,4},{5,6}}
	fmt.Println(array)
	var array2=[...][2]int{{1,2},{3,4},{5,6},{7,8}}
	//var array3=[4][...]int{{1,2},{3,4},{5,6},{7,8}}
	fmt.Println(array2)
	fmt.Println(len(array2))
}
