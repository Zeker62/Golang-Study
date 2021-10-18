package main
//多敲两遍
import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//	初始化随机种子
	rand.Seed(time.Now().UnixNano())

	var maps=make(map[string]int,200)
	for i:=0;i<50;i++{
		key:=fmt.Sprintf("student:%02d",i)
		value:=rand.Intn(200)
		maps[key]=value
	}
	// 将key存入到keys切片当中去
	keys:=make([]string,0,200)
	for key:=range maps{
		keys=append(keys,key)
	}
	fmt.Println(keys)

	sort.Strings(keys)
	for _,key:=range keys{
		fmt.Println(key,"=",maps[key])
	}

}
