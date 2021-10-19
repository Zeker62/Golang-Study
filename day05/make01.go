package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var l map[string]int32
	fmt.Printf("l的地址是：%p\n",&l)
	fmt.Printf("l的大小是：%v\n",unsafe.Sizeof(l))
	l=make(map[string]int32,10)
	l["aaa"]=122
	fmt.Printf("l[aaa]的大小是：%v",unsafe.Sizeof(l["aaa"]))
}
