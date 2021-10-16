package main

import "fmt"

func main() {
	var array2=[...][2]int{{1,2},{3,4},{5,6},{7,8}}
	for i:=0;i<len(array2);i++{
		for _,a:=range array2[i]{
			fmt.Print(a,"\t")
		}
		fmt.Println()
	}
}
//1	2
//3	4
//5	6
//7	8