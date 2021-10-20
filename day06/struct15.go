package main

import "fmt"

type person15 struct {
	name string
	age int
	infor15
}
type infor15 struct {
	address string
	telephone int64
}

func main() {

	var p person15
	p.infor15.telephone=34567890
	p.name="甄嬛"
	p.address="紫阳小街" //匿名字段名可以省略
	fmt.Printf("The information of p is :%#v",p)
}
