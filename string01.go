package main

import "fmt"

func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1) //先转为byte类型
	byteS1[0] = 'p'// 修改
	fmt.Println(string(byteS1)) //在转为string类型

	s2 := "白萝卜"
	runeS2 := []rune(s2) //先转为rune类型
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
func main()  {
	changeString()
}
