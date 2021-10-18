package main

import "fmt"

func sum009(x int,y int)int{
	return x+y
}



func summary009(x int,y int,function func( int, int)int)int{
	return function(x,y)
}


func main() {
	result:=summary009(3,5,sum009)
	fmt.Println(result)
}
