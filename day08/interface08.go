package main

import "fmt"

type animals08 interface {
	cat08
	dog08
}
type cat08 interface {
	say()
}
type dog08 interface {
	move()
}
type cat08_ struct {
	name string
}
func (c cat08_) say(){
	fmt.Println(c.name,"在叫")
}
func (c cat08_) move(){
	fmt.Println(c.name,"会动")
}
func main() {
	var a animals08
	var c=cat08_{
		name:"Anim",
	}
	a=c
	a.say()
	a.move()

}
//Anim 在叫
//Anim 会动

