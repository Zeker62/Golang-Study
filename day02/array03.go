package main

import "fmt"

func main() {
	var array=[7]int{1:2,5:100,2:10}
	var array2=[...]int{1:2,20:10}
	fmt.Print(array," ")
	fmt.Printf("type of array2:%T \n",array)
	fmt.Print(array2," ")
	fmt.Printf("type of array2:%T",array2)
}
//[0 2 10 0 0 100 0] type of array2:[7]int
//[0 2 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 10] type of array2:[21]int