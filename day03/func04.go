package main

import "fmt"

func ca002(x int,y int)(sum ,sub int){
	sum=x+y
	sub=x-y
	return
}
func main() {
	fmt.Println(ca001(5,7)) //12 -2
}
