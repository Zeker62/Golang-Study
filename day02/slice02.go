package main

import "fmt"

func main() {
	var array=[6]int{1,2,3,4,5,6}
	s:=array[:3]
	fmt.Printf("s:%v\t length:%v \tcap:%v",s,len(s),cap(s))
	fmt.Printf("\n%v",array[:3])
}
