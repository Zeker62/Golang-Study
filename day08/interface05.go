package main

import "fmt"

type saying05 interface {
	say()
}
type name05 interface {
	tell()
}
type dog05 struct {
	name string
}
func (d dog05) say(){
	fmt.Println(d.name,"会说话")
}
func (d dog05) tell(){
	fmt.Println(d.name,"有名字")
}

func main() {
	var s saying05
	var n name05
	var d=dog05{
		name:"Bob",
	}
	s=d
	n=d
	s.say()
	n.tell()

}
