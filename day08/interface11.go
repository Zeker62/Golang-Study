package main

import "fmt"

func main() {
	var a =make(map[interface{}]interface{})
	a[123]="key为数字类型的map"
	a["hello"]=1.223
	a[3.445]=true
	b:=100
	a[&b]=*&b
	for key,value:=range a{
		fmt.Printf("key type is :%T\t\t\tkey is %v\t\t\tvalue is:%v\n",key,key,value)
	}

}
