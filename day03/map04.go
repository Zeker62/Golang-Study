package main

import "fmt"

func display(maps map[string]int){
	for a,b:=range maps{
		fmt.Println("maps[",a,"]",b)
	}
}
func main() {
	maps:=make(map[string]int,10)
	maps["bbb"]=345
	maps["ccc"]=678
	maps["aaa"]=123
	maps["ddd"]=0

	fmt.Println("删除前")
	display(maps)
	fmt.Println("删除ccc")
	delete(maps,"ccc")
	display(maps)
}
//maps[ bbb ] 345
//maps[ ccc ] 678
//maps[ aaa ] 123
//maps[ ddd ] 0
//删除ccc
//maps[ aaa ] 123
//maps[ ddd ] 0
//maps[ bbb ] 345