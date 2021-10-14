// iota是go语言的常量计数器，只能在常量的表达式中使用。

// iota在const关键字出现时将被重置为0。
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
// 使用iota能简化定义，在定义枚举时很有用。
package main

import "fmt"

const (
	a1 = iota
	a2
	a3
	_ //使用_跳过某些值
	_
	a4
	a5 = 100
	a6
	a7 = iota //iota声明中间插队
	a8

	// 0 1 2 5 100 100 8 9
)

// 定义数量级
//这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。
//同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
	// 1024 1048576 1073741824 1099511627776 1125899906842624
)

// 多个iota定义在一行
const (
	a, b = iota, iota         //0,0
	c, d = iota, iota         // 1,1
	e, f = iota + 1, iota + 2 //3,4

	// 0 0 1 1 3 4
)

func main() {
	fmt.Println(a1, a2, a3, a4, a5, a6, a7, a8)
	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a, b, c, d, e, f)
}
