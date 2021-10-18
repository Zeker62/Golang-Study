package main

import "fmt"

func sum002(x int,y...int)int{
	sum:=x
	for _,v :=range y{
		sum+=v
	}
	return sum
}
func main() {
	fmt.Println(sum002(1,2,3,4,5))
}
