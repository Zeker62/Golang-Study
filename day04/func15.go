package main

import "fmt"


func calculate15(x int)(func (int)int,func(int)int){
	add:=func (y int)int{
		return x+y
	}
	sub:=func(y int)int{
		return x-y
	}
	return add,sub
}

func main() {
	var f1,f2=calculate15(10)
	fmt.Println(f1(5),f2(10)) //15 0
	fmt.Println(f1(10),f2(20)) //20 -10
}
