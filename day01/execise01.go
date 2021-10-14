package main
//编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。

import (
	"fmt"
)
var(
	aaa int
	bbb float64
	ccc bool
	ddd string
)
func main() {
	aaa,bbb,ccc,ddd=3,3.2,true,"aaa"
	// 只有printf才能使用占位符
	fmt.Printf("aaa=%T\nbbb=%T\nccc=%T\nddd=%T", aaa, bbb, ccc, ddd)
}
