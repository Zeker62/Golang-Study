package main

import "fmt"

type saying04 interface {
	say()
}
type dog04 struct {}

func (d *dog04) say(){ //指针类型的接收者
	fmt.Println("wang wang wnag")
}
func main() {
	var animal saying04
	// var WWW=dog04{} 报错不能使用
	// animal=WWW
	//animal.say()
	var SSS=&dog04{}
	animal=SSS //等同于animal=*SSS
	animal.say()
}
