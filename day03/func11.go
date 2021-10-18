package main

import "fmt"

func main() {
	sum:=func(x,y int){
		fmt.Println(x+y)
	}
	sum(10,20)
	fmt.Println(func(x,y int)int{
		return x+y+y
	}(10,20))

}
