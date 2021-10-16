package main

import "fmt"

func main() {
	var array[]int
	array=append(array,1,2,3,4)
	fmt.Println(array)
	array2:=[]int{5,6,7,8}
	//array3:=[...]int{10,11,12}
	array=append(array,array2...)
	fmt.Println(array)
	//array=append(array,array3)
	//fmt.Println(array)

}
