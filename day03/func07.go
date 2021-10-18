package main

import "fmt"

func main() {
	a:=[]int{1,2,3,4,5,6,7}
	for i:=0;i<len(a);i++{
		b:=i
		if b<10{
			c:=10
			fmt.Println(b,c)
		}
	}
	//fmt.Println(i)
	//fmt.Println(b)
	//fmt.Println(c)
}
