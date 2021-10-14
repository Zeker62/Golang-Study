// Go语言支持两种浮点型数：float32和float64。这两种浮点型数据格式遵循IEEE 754标准：
// float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。
//  float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。

// 打印浮点数时，可以使用fmt包配合动词%f，代码如下：

package main

import (
	"fmt"
	"math"
)

// complex64和complex128
// 复数
var c1 complex64 = 1 + 2i
var c2 complex128 = 2 + 3i

// Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。
// 注意：
// 布尔类型变量的默认值为false。
// Go 语言中不允许将整型强制转换为布尔型.
// 布尔型无法参与数值运算，也无法与其他类型进行转换。

func main() {
	fmt.Println(math.MaxFloat64)

	// 浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// 复数
	fmt.Println(c1)
	fmt.Println(c2)

	// 	3.141593
	// 3.14
	// (1+2i)
	// (2+3i)

}
