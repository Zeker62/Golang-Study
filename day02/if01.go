
package main

import "fmt"

func main(){
	//一般写法
	var your_age=18
	if your_age<18{
		fmt.Println("you are too young!")
	}else if 18<=your_age && your_age<35{
		fmt.Println("it is the best")
	}else{
		fmt.Println("you are too old to do this job")
	}
	fmt.Println("your age is",your_age)
}


