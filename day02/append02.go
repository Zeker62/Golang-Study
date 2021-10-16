package main

import "fmt"

func main() {
	var array[]int
	for i:=0;i<10;i++{
		array=append(array,i)
		fmt.Printf("%v len:%d cap:%d addr:%p \n",array,len(array),cap(array),array)
	}
}
//[0] len:1 cap:1 addr:0xc0000aa058
//[0 1] len:2 cap:2 addr:0xc0000aa0a0
//[0 1 2] len:3 cap:4 addr:0xc0000a80a0
//[0 1 2 3] len:4 cap:4 addr:0xc0000a80a0
//[0 1 2 3 4] len:5 cap:8 addr:0xc0000c20c0
//[0 1 2 3 4 5] len:6 cap:8 addr:0xc0000c20c0
//[0 1 2 3 4 5 6] len:7 cap:8 addr:0xc0000c20c0
//[0 1 2 3 4 5 6 7] len:8 cap:8 addr:0xc0000c20c0
//[0 1 2 3 4 5 6 7 8] len:9 cap:16 addr:0xc0000d4080
//[0 1 2 3 4 5 6 7 8 9] len:10 cap:16 addr:0xc0000d4080