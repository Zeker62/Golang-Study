//Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。
//Go 语言里的字符串的内部实现使用UTF-8编码。
//字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符

// Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

// 转义符	含义
// \r	回车符（返回行首）
// \n	换行符（直接跳到下一行的同列位置）
// \t	制表符
// \'	单引号
// \"	双引号
// \\	反斜杠
package main

import "fmt"

func main() {
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	// Go语言中要定义一个多行字符串时，就必须使用反引号字符：
	s1 := `第一行
第二行
第三行
`
	fmt.Println(s1)
	// len(str)	求长度
	// +或fmt.Sprintf	拼接字符串
	// strings.Split	分割
	// strings.contains	判断是否包含
	// strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
	// strings.Index(),strings.LastIndex()	子串出现的位置
	// strings.Join(a[]string, sep string)	join操作

}
