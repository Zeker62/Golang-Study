package main

import "fmt"

type saying07 interface {
	say()
	move()
}

type dog07 struct {
	name string
}

func (d dog07) move() {
	panic("implement me")
}

func (d dog07) say(){
	fmt.Println(d.name,"又要开始叫了")
}
type cat07 struct {
	dog07
}

func (c cat07) move()  {
	fmt.Println(c.name,"又开始动了")
}


func main() {
	var s saying07
	d:=dog07{
		name:"Toms",
	}
	c:=cat07{
		d,

	}
	s=d
	s.say()
	s=c
	s.move()
}
