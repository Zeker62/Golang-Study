package main
//导入语句
import "fmt"

//函数外面只能放置标识符的声明，（变量声明、常量声明啥的）
//fmt.println("sss")


//程序的入口
func main() {
	fmt.Println("Hello World")
}
// 使用 go build  main.go 可以生成二进制文件
// 使用 go run main.go 可以直接执行main.go文件
// 使用 main.exe 执行二进制文件
