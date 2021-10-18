package main

import "fmt"

func isOK(ok bool,v int){
	if ok{
		fmt.Println(v,ok)
	}else{
		fmt.Println(v,ok)
		fmt.Println("没有此键")
	}
}
func main() {
	name:=make(map[string]int,8)
	name["aaa"]=123
	name["bbb"]=456
	v,ok:=name["ccc"]
	isOK(ok,v)
	v,ok=name["aaa"]
	isOK(ok,v)
}
//0 false
//没有此键
//123 true