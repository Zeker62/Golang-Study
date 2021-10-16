package main

import "fmt"

func main() {
	array:=make([]int,3)
	s2:=array
	s2[2]=100
	fmt.Print(s2,"\n")
	fmt.Print(array)
}
