package main

import "fmt"

func main() {
	var array=[]int{1,2,3,4,4,5,6,7,8}
	array2:=array //浅复制，指针
	array2[3]=100
	fmt.Println(array)
	fmt.Println(array2)

	array3:=make([]int,8,8)
	copy(array3,array) // 复制，开辟内存空间传递值
	array3[5]=1000
	fmt.Println(array3)
	fmt.Println(array)
}
//[1 2 3 100 4 5 6 7 8]
//[1 2 3 100 4 5 6 7 8]
//[1 2 3 100 4 1000 6 7]
//[1 2 3 100 4 5 6 7 8]