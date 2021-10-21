package main

import (
	"fmt"
	c "github.com/Zeker62/day07/package01/calc"
)


func main() {
	fmt.Println("请输入x和y,用空格分隔")
	var x int
	var y int
	fmt.Scanf("%d %d",&x,&y)
	fmt.Println(c.Sub(x,y),c.Mut(x,y),c.Sum(x,y),c.Dec(x,y))
}
