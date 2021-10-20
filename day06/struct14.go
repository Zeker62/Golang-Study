package main

import "fmt"

type person14 struct {
	name string
	age int
	infor infor14
}
type infor14 struct {
	address string
	telephone int64
}

func main() {

	p:=person14{
		name:"Tom",
		age:33,
		infor:infor14{
			address:"Fire Street",
			telephone: 1228847389,
		},
	}
	fmt.Printf("The information of p is :%#v",p)
}
