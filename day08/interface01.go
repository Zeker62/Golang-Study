package main

import "fmt"

type say interface{
	saying()
}
type dog struct{

}
type cat struct {

}
func (d dog)say(){
	fmt.Println("wang wang wang")
}
func (c cat)say(){
	fmt.Println("miao miao miao ")
}
func main() {
	var d dog
	var c cat
	d.say()
	c.say()
}
