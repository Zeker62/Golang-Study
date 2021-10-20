package main

import "fmt"

type person11 struct{
	name string
	age int
}
//构造函数
func newPerson11(name string,age int) *person11{
	return &person11{
		name:name,
		age:age,
	}
}
//方法
func (p *person11) change_age(new_age int){
	p.age=new_age
}
func (p person11) tell11(){
	fmt.Printf("%s is %v",p.name,p.age)
}
func (p person11) change_name(new_name string){
	p.name=new_name
	//fmt.Println(p.name)
}
func main() {
	var p11=newPerson11("小红",22)
	p11.change_age(18)
	p11.change_name("小王")
	p11.tell11()

}
