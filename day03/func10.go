package main

import "fmt"

func sum010(x int,y int)(func (int,int)int){
	s:=x+y
	if s>100 {
		return sum010(s, y)
	}
	return nil
}


func main() {
	fmt.Println(sum010(5,3))
}
