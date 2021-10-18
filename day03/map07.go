package main

import "fmt"

func main() {
	maps:=make(map[string][]string,5)
	fmt.Println(maps)
	fmt.Println("初始化……")
	key:="City"
	value,ok :=maps[key]
	if !ok{
		value=make([]string,0,6)
	}
	value=append(value,"Beijing","Shanghai")
	maps[key]=value
	fmt.Println(maps)

}
//map[]
//初始化……
//map[City:[Beijing Shanghai]]
