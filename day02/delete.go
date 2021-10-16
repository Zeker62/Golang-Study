package main

import "fmt"

func main() {
	var array=[]int{1,2,3,4,5,6,7}
	fmt.Println(array)
	array=append(array[:2],array[3:]...)
	fmt.Println(array)

}
//[1 2 3 4 5 6 7]
//[1 2 4 5 6 7]