// 数字字面量语法（Number literals syntax）
// Go1.13版本之后引入了数字字面量语法，这样便于开发者以二进制、八进制或十六进制浮点数的格式定义数字，例如：

// v := 0b00101101， 代表二进制的 101101，相当于十进制的 45。
// v := 0o377，代表八进制的 377，相当于十进制的 255。
// v := 0x1p-2，代表十六进制的 1 除以 2²，也就是 0.25。

// 而且还允许我们用 _ 来分隔数字，比如说： v := 123_456 表示 v 的值等于 123456。

// 我们可以借助fmt函数来将一个整数以不同进制形式展示。

package main

import "fmt"

//go语言中无法定义二进制数
func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10 %d表示整数，%s表示的是字符串
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头，因为八进制是三位二进制，就是作为32位系统的时候前八位位0
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	// var d int = 18
	// fmt.Println("%x \n", d) //输出%x\n 18

	// 查看变量类型
	fmt.Printf("%T\n", a)

	var e int8 = 9
	fmt.Println("%T\n",e)
}
