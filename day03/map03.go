package main

import "fmt"

func main() {
	maps:=make(map[string]int,10)
	maps["bbb"]=345
	maps["ccc"]=678
	maps["aaa"]=123

	for a,b:=range maps{
		fmt.Println("maps[",a,"]",b)
	}
}
