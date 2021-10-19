package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer fmt.Print(1)
	defer fmt.Print(2)
	defer fmt.Print(3)
	fmt.Println("End")
}
//Start
//End
//321