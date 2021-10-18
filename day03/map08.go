package main

import "fmt"

//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
func main() {
	s:="how do you do"
	array:=[]rune(s)
	maps:=make(map[string]int,10)
	for _,c:=range array{
		value,ok :=maps[string(c)]
		if ok{
			value+=1
			maps[string(c)]=value
		}else{
			maps[string(c)]=1
		}
	}
	for keys,values:=range maps{
		if keys!=" "{
			fmt.Printf("%v=%d ",keys,values)
		}

	}
}
