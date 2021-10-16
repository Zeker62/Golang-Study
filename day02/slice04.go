package main

import "fmt"

func main() {
	var array=[8]int{1,2,3,4,5,6,7,8}
	fmt.Printf("array: %v\t array[:3:5]:%v\t len(array[:3:5]):%v\t cap(array[:3:4]):%v\t",array,array[:3:5],len(array[:3:5]),cap(array[:3:4]))

}
