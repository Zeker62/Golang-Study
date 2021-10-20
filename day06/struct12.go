package main

import "fmt"

type MyInt int64

func (i MyInt) sum12(a,b MyInt)int64{
	i=a+b
	return int64(i)
}
func main() {
	var i11 MyInt
	c:=i11.sum12(10,20)
	fmt.Printf("数据类型是i11的对象调用方法sum12的结果是：%v",c)
}
