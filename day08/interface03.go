package main

import "fmt"

type saying03 interface {
	say()
}
type dog03 struct {}

func (d dog03) say(){
	fmt.Println("wang wang wnag")
}
func main() {
	var animal saying03
	var WWW=dog03{}
	animal=WWW
	animal.say()
	var SSS=&dog03{}
	animal=SSS //等同于animal=*SSS
	animal.say()
}
