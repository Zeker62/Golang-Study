// 匿名变量

//在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示
package main
import "fmt"
func foo() (int, string) {
	return 10, "aaa"
}
func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}

// 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。 (在Lua等编程语言里，匿名变量也被叫做哑元变量。)

// 注意事项：

// 函数外的每个语句都必须以关键字开始（var、const、func等）
// :=不能使用在函数外。
// _多用于占位，表示忽略值。