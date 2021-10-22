package main

import "fmt"

type say02 interface {
	say()
}
type dog02 struct{}
type cat02 struct{}
func (d dog02)say(){
	fmt.Println("wang wang wang")
}
func (c cat02)say(){
	fmt.Println("miao miao miao ")
}
func main() {
	var s say02
	d:=dog02{}
	c:=cat02{}
	s=d
	s.say()
	s=c
	s.say()
}
