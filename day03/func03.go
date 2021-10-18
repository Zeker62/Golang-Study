package main

import "fmt"

func ca001(x int,y int)(int,int){
	return x+y,x-y
}
func main() {
	fmt.Println(ca001(5,7)) //12 -2
}
