package main

import (
	"fmt"
	"io"
	"os"
)

// 定义日志写入接口
type write_logger interface {
	write(string) 				//定义一个写入日志文件的方法
}
// 日志结构体
type logger struct{
	filename string
}

// 判断当前日志文件是否存在
func check_file_is_exist(filename string)bool{
	_,err:=os.Stat(filename)	//os.Stat是获取文件属性的方法，如果该文件不存在会报错，将值传入err
	if os.IsNotExist(err){		//判断这个是不是文件不存在的异常错误
		return false
	}
	return true					//文件存在就会返回true
}

// 文件写入接口
func (f1 *logger) write(message string){	//write是接口write_logger的一个方法
	var file *os.File						//定义指针类型的文件，这样可以篡改
	var err  error							//源码可以看见，error的类型本质上是string
	if check_file_is_exist(f1.filename){	//看是否存在这个文件
		file,err=os.OpenFile(f1.filename,os.O_APPEND|os.O_WRONLY,0666)
		// 此时文件被打开，只写模式，并且追加内容，权限相当于Linux的666
		fmt.Println("文件存在")
	}else{
		file,err = os.Create(f1.filename)
		fmt.Println("文件不存在，已创建")
	}
	defer file.Close()
	n,err:=io.WriteString(file,message+"\n")  //写入message，写入 成功就返回字节数
	if err!=nil{
		panic(err)
	}
	fmt.Println("写入了%d个字节",n)

}
type ConsoleLogger struct{} //空结构体对象


func (c1 *ConsoleLogger) write(message string){
	fmt.Println(message) //ConsoleLogger对象实现的就是这个方法
}

func homework(){
	var log write_logger
	filelogger:=&logger{"log.txt"}
	log=filelogger
	var s string
	fmt.Scanf("%s",&s)
	log.write(s)
	log.write("how are you")

	consoleLogger:=&ConsoleLogger{}
	log=consoleLogger
	log.write("Heo")
	log.write("how are you")
}
func main(){
	homework()
}