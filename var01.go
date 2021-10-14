package main
import "fmt"
var name string
////批量声明
var(
age int
isOk bool
)
func main(){
	name="a" //报错：rune字符过多，只能用双引号不能用单引号
	age=12
	isOk=true
	// go语言中必须使用，不使用就无法编译
	fmt.Printf("name:%s",name) //%s是个占位符，用name这个变量去替换
	fmt.Println(age) //打印完内容自动换行
	fmt.Print(isOk)
	fmt.Println()
	// 快捷变量声明，可以自动识别函数类型
	var s2="20"
	//简短变量声明：只能在函数中使用，可以自动识别函数类型
	s3:="hahah"
	fmt.Println(s2,s3)

	// var s4报错，必须要指明数据类型
	// s4='aaa'
	//fmt.Println(s4)

}
