package main

import "fmt"

type move06 interface {
	move()
}
type dog06 struct {
	name string
}
type cat06 struct {
	name string
}
func (d dog06) move(){
	fmt.Println(d.name,"会动")
}
func (c cat06) move(){
	fmt.Println(c.name,"会动")
}
func main() {
	var m move06
	var d=dog06{
		name:"Tom",
	}
	var c=cat06{
		name:"Crazy",
	}
	m=d
	m.move()
	m=c
	m.move()
}
//Tom 会动
//Crazy 会动