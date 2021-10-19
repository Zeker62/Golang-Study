package main

import "fmt"

func main() {
	var a *int
	a=new(int)
	*a=10
	fmt.Printf("a的地址为 %p \na的值为 %v\n*a的地址为 %p\n*a的值为 %v",&a,a,&*a,*a)
}
