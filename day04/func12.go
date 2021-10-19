package main

import "fmt"

func sum012(x,y int)func()int{
	return func() int {
		return x+y
	}
}

func main() {
	f:=sum012(5,6)
	fmt.Println(f())
	fmt.Println(sum012(7,8)())
}
