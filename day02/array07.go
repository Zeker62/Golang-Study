package main

import "fmt"

func change(a [3]int)  {
	a[2]=22
	fmt.Println("在函数体内部时",a)
}
func main() {
	var array=[3]int{1,2,3}
	change(array)
	fmt.Println("最终修改后：",array)

}
//在函数体内部时 [1 2 22]
//最终修改后： [1 2 3]