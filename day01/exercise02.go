package main

import "fmt"

//编写代码统计出字符串"hello沙河小王子"中汉字的数量。

func main(){
	s:="hello沙河小王子"
	count:=0
	for _,r :=range s{
		if int(r)>1000{
			count+=1
			fmt.Printf("%v(%c)\n",r,r)
		}
	}
	fmt.Println(count)
}
