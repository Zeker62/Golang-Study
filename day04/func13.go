package main

import "fmt"

func sum013(x int)func(int)int{
	return func(y int)int{
		x+=y
		return x
	}
}

func main() {
	var f=sum013(30)
	fmt.Println(f(20)) //50
	fmt.Println(f(30)) //80
	fmt.Println(f(40)) //120
	var s = sum013(10)
	fmt.Println(s(10)) //20
}
