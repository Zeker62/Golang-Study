package main

import "fmt"

func sum001(a int,b int)int{
	fmt.Println("求和函数")
	sum:=a+b
	return sum
}

func main() {
	fmt.Println(sum001(3,5))
}
