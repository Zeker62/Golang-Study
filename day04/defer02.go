package main

import "fmt"

func test_defer()string{
	fmt.Println("这是没有defer的语句")
	defer fmt.Println("这是defer语句")
	return "这是return"
}
func main() {
	fmt.Println(test_defer())
}
