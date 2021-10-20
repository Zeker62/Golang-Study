package main

import "fmt"

type Int8 int8
type Int64=int64
type Person13 struct{
	int
	int64
	int8
	Int8
	Int64
}
func main() {
	p1:=Person13{
		1,2,3,4,5,
	}
	fmt.Printf("value of p1: %#v",p1)

}
