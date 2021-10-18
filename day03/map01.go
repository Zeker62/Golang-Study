package main

import "fmt"

func main() {
	name:=make(map[string]int)
	name["张三"]=23
	name["李四"]=25
	fmt.Println(name)
	fmt.Println(name["李四"])
	fmt.Printf("type of map:%T\n",name)

	address:=map[string]string{
		"Tom":"break",
		"Alice":"continue",
	}
	fmt.Println(address)

	//age:=map[string]string
	//age["Cab"]="aaa"
	//fmt.Println(age)

}
