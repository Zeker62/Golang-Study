package main

import "fmt"
func First_recover(){
	fmt.Println("Recover的使用")
}
func A_recover(){
	defer func() {
		error := recover()
		if error != nil {
			fmt.Println("程序出现了错误，现已恢复")
		}
	}()
	panic("程序有错")
}

func main() {
	First_recover()
	A_recover()
}
