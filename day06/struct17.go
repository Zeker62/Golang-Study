package main

import "fmt"
// 继承
type Animal17 struct {
	name string
}
func (a *Animal17) move(){
	fmt.Printf("%s 会动\n",a.name)
}
type Chicken17 struct {
	food string
	*Animal17 //通过匿名结构体的嵌套实现了继承
}
// 定义方法
func (a *Chicken17) eat(){
	fmt.Printf("%s 会吃 %s\n",a.name,a.food)
}
func main() {
	a:=&Chicken17{
		food :"粮食",
		Animal17:&Animal17{
			name:"鸭子",
		},
	}
	a.move()
	a.eat()
}
