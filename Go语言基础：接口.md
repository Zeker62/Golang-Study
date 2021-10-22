

接口（interface）定义了一个对象的行为规范，**只定义规范不实现**，由具体的对象来实现规范的细节。

# 接口

## 接口类型

在Go语言中接口（interface）是一种类型，一种抽象的类型。

`interface`是一组`method`的集合，是`duck-type programming`的一种体现。接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。

为了保护你的Go语言职业生涯，请牢记接口（interface）是一种类型。

## 为什么要使用接口

```go
type Cat struct{}

func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }

func main() {
	c := Cat{}
	fmt.Println("猫:", c.Say())
	d := Dog{}
	fmt.Println("狗:", d.Say())
}
```

上面的代码中定义了猫和狗，然后它们都会叫，你会发现main函数中明显有重复的代码，如果我们后续再加上猪、青蛙等动物的话，我们的代码还会一直重复下去。那我们能不能把它们当成“能叫的动物”来处理呢？

像类似的例子在我们编程过程中会经常遇到：

- 比如一个网上商城可能使用支付宝、微信、银联等方式去在线支付，我们能不能把它们当成“支付方式”来处理呢？

- 比如三角形，四边形，圆形都能计算周长和面积，我们能不能把它们当成“图形”来处理呢？

- 比如销售、行政、程序员都能计算月薪，我们能不能把他们当成“员工”来处理呢？

Go语言中为了解决类似上面的问题，就设计了接口这个概念。接口区别于我们之前所有的具体类型，接口是一种抽象的类型。当你看到一个接口类型的值时，你不知道它是什么，唯一知道的是通过它的方法能做什么。

## 接口的定义

Go语言提倡面向接口编程。

每个接口由数个方法组成，接口的定义格式如下：

```go
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

其中：

- 接口名：使用`type`将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加`er`，如有写操作的接口叫`Writer`，有字符串功能的接口叫`Stringer`等。接口名最好要能突出该接口的类型含义。
- 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
- 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

举个例子：

```go
type writer interface{
    Write([]byte) error
}
```

当你看到这个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的Write方法来做一些事情。

## 实现接口的条件

一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个**需要实现的方法列表**。

如下面的例子：

```go
package main

import "fmt"

type say interface{
   saying()
}
type dog struct{}
type cat struct {}
func (d dog)say(){
   fmt.Println("wang wang wang")
}
func (c cat)say(){
   fmt.Println("miao miao miao ")
}
func main() {
   var d dog
   var c cat
   d.say()
   c.say()
}
//wang wang wang
//miao miao miao 
```

## 接口类型变量

那实现了接口有什么用呢？

接口类型变量能够存储所有实现了该接口的实例。 例如上面的示例中，`Sayer`类型的变量能够存储`dog`和`cat`类型的变量。

```go
package main

import "fmt"

type say02 interface {
   say()
}
type dog02 struct{}
type cat02 struct{}
func (d dog02)say(){
   fmt.Println("wang wang wang")
}
func (c cat02)say(){
   fmt.Println("miao miao miao ")
}
func main() {
   var s say02
   d:=dog02{}
   c:=cat02{}
   s=d
   s.say()
   s=c
   s.say()
}
```

## 值接收者和指针接收者实现接口的区别

使用值接收者实现接口和使用指针接收者实现接口有什么区别呢？接下来我们通过一个例子看一下其中的区别。

我们有一个`saying03`接口和一个`dog`03结构体。

### 值接收者实现接口

```go
package main

import "fmt"

type saying03 interface {
   say()
}
type dog03 struct {}

func (d dog03) say(){ //值接收者
   fmt.Println("wang wang wnag")
}
func main() {
   var animal saying03
   var WWW=dog03{}
   animal=WWW
   animal.say()
   var SSS=&dog03{}
   animal=SSS //等同于animal=*SSS
   animal.say()
}
//wang wang wnag
//wang wang wnag
```

从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。系统会将指针类型的变量自动求值，增加了代码的容错率

### 指针接收者实现接口

同样的代码我们再来测试一下使用指针接收者有什么区别：

```go
package main

import "fmt"

type saying04 interface {
   say()
}
type dog04 struct {}

func (d *dog04) say(){ //指针类型的接收者
   fmt.Println("wang wang wnag")
}
func main() {
   var animal saying04
   // var WWW=dog04{} 报错不能使用
   // animal=WWW
   //animal.say()
   var SSS=&dog04{}
   animal=SSS //等同于animal=*SSS
   animal.say()
}
```



所以，接收者如果是值，遇到指针类型会自动转换为值，但接收者如果是指针，只能接收指针类型结构体咯

### 面试题

下面的代码可不可以通过编译？

```go
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = Student{} //此处出现编译错误，应当传递指针
    // 修正为：var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
```

## 类型与接口的关系

### 一个类型实现多个接口

一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。

```go
package main

import "fmt"

type saying05 interface {
   say()
}
type name05 interface {
   tell()
}
type dog05 struct {
   name string
}
func (d dog05) say(){
   fmt.Println(d.name,"会说话")
}
func (d dog05) tell(){
   fmt.Println(d.name,"有名字")
}

func main() {
   var s saying05 //接口1
   var n name05 //接口2
   var d=dog05{
      name:"Bob",
   }
    //一个对象，两个接口
   s=d
   n=d
   s.say()
   n.tell()

}
//Bob 会说话
//Bob 有名字
```

### 多个类型实现同一接口

Go语言中不同的类型还可以实现同一接口

```go
package main

import "fmt"

type move06 interface {
   move()
}
type dog06 struct {
   name string
}
type cat06 struct {
   name string
}
func (d dog06) move(){
   fmt.Println(d.name,"会动")
}
func (c cat06) move(){
   fmt.Println(c.name,"会动")
}
func main() {
   var m move06
   var d=dog06{
      name:"Tom",
   }
   var c=cat06{
      name:"Crazy",
   }
   m=d
   m.move()
   m=c
   m.move()
}
//Tom 会动
//Crazy 会动
```

并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

```go
package main

import "fmt"

type saying07 interface {
   say()
   move()
}

type dog07 struct {
   name string
}

func (d dog07) move() {
   panic("implement me")
}

func (d dog07) say(){
   fmt.Println(d.name,"又要开始叫了")
}
type cat07 struct {
   dog07
}

func (c cat07) move()  {
   fmt.Println(c.name,"又开始动了")
}


func main() {
   var s saying07
   d:=dog07{
      name:"Toms",
   }
   c:=cat07{d}
   s=d
   s.say()
   s=c
   s.move()
}
```

## 接口嵌套

接口与接口间可以通过嵌套创造出新的接口。

嵌套得到的接口的使用与普通接口一样。

```go
package main

import "fmt"

type animals08 interface {
   cat08
   dog08
}
type cat08 interface {
   say()
}
type dog08 interface {
   move()
}
type cat08_ struct {
   name string
}
func (c cat08_) say(){
   fmt.Println(c.name,"在叫")
}
func (c cat08_) move(){
   fmt.Println(c.name,"会动")
}
func main() {
   var a animals08
   var c=cat08_{
      name:"Anim",
   }
   a=c
   a.say()
   a.move()

}
//Anim 在叫
//Anim 会动
```

## 空接口

### 空接口的定义

空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。

空接口类型的变量可以存储任意类型的变量。正是由于这个特性就有很多的功能

```go
package main

import "fmt"

func main() {
   var inter interface{}
   s:="这是一个字符串"
   inter=s
   fmt.Println(inter)// 字符串类型
   a:=10
   inter=a
   fmt.Println(inter) //int类型
   c:=100.0002
   inter=c
   fmt.Println(inter)//float类型
}

//这是一个字符串
//10
//100.0002
```

### 空接口的应用

#### 空接口作为函数的参数

使用空接口实现可以接收任意类型的函数参数。

```go
package main

import "fmt"

func show_data_10(a interface{}){
	fmt.Printf("通过interface接收了值为%v ，%T类型的变量\n",a,a)
}

func main() {
	var a interface{}
	b:=1000
	a=b
	show_data_10(a)
	c:=&b //指针类型
	a=c
	show_data_10(a)
	d:="字符串"
	a=d
	show_data_10(a)
}
```

效果：

```
通过interface接收了值为1000 ，int类型的变量
通过interface接收了值为0xc00000a0a8 ，*int类型的变量
通过interface接收了值为字符串 ，string类型的变量
```

#### 空接口作为map的值

使用空接口实现可以保存任意值的字典。

```go
package main

import "fmt"

func main() {
	var a =make(map[interface{}]interface{})
	a[123]="key为数字类型的map"
	a["hello"]=1.223
	a[3.445]=true
	b:=100
	a[&b]=*&b
	for key,value:=range a{
		fmt.Printf("key type is :%T\t\t\tkey is %v\t\t\tvalue is:%v\n",key,key,value)
	}

}
```

效果：

```ABAP
key type is :*int			key is 0xc0000aa058			value is:100
key type is :int			key is 123			value is:key为数字类型的map
key type is :string			key is hello			value is:1.223
key type is :float64			key is 3.445			value is:true
```

## 类型断言

空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？

### 接口值

一个接口的值（简称接口值）是由`一个具体类型`和`具体类型的值`两部分组成的。这两部分分别称为接口的`动态类型`和`动态值`。

我们来看一个具体的例子：

```go
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil
```

请看下图分解：![接口值图解](https://www.liwenzhou.com/images/Go/interface/interface.png)

想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：

```go
x.(T)
```

其中：

- x：表示类型为`interface{}`的变量
- T：表示断言`x`可能是的类型。

该语法返回两个参数，第一个参数是`x`转化为`T`类型后的变量，第二个值是一个布尔值，若为`true`则表示断言成功，为`false`则表示断言失败。

例如：

```go
package main

import "fmt"

func main() {
   var a interface{}
   a="aabbcc"
   v,ok:=a.(string)
   if ok{
      fmt.Println(v)
   }else{
      fmt.Println("断言错误")
   }

}
```

如果需要多次断言的话，还可以创建一个断言函数：

```go
package main

import "fmt"

func switchtype(a interface{}){
   switch v:=a.(type){
   case string:
      fmt.Printf("x is a string，value is %v\n", v)
   case int:
      fmt.Printf("x is a int is %v\n", v)
   case bool:
      fmt.Printf("x is a bool is %v\n", v)
   default:
      fmt.Println("unsupport type！")
   }
}

func main() {
   var a interface{}
   a="aavvss"
   switchtype(a)
}
//x is a string，value is aavvss
```

因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛。

关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。





## 练习：

使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库。

### os.Stat的用法：

```go
package main

import (
    "fmt"
    "os"
)

func main() {
	fileInfo, err := os.Stat(`C:\Users\Administrator\Desktop\应用商店.txt`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileInfo.Name())    //应用商店.txt
	fmt.Println(fileInfo.IsDir())   //false  判断是否是目录
	fmt.Println(fileInfo.ModTime()) //2019-12-05 16:59:36.8832788 +0800 CST   文件的修改时间
	fmt.Println(fileInfo.Size())    //3097  文件大小
    fmt.Println(fileInfo.Mode())    // -rw-rw-rw-  读写属性
	fmt.Println(fileInfo.Sys())     //&{32 {2160608986 30778972} {2160608986 30778972} {1375605524 30780234} 0 3097}
}
```

### 答案：

```go
package main

import (
   "fmt"
   "io"
   "os"
)

// 定义日志写入接口
type write_logger interface {
   write(string)           //定义一个写入日志文件的方法
}
// 日志结构体
type logger struct{
   filename string
}

// 判断当前日志文件是否存在
func check_file_is_exist(filename string)bool{
   _,err:=os.Stat(filename)   //os.Stat是获取文件属性的方法，如果该文件不存在会报错，将值传入err
   if os.IsNotExist(err){    //判断这个是不是文件不存在的异常错误
      return false
   }
   return true                //文件存在就会返回true
}

// 文件写入接口
func (f1 *logger) write(message string){   //write是接口write_logger的一个方法
   var file *os.File                 //定义指针类型的文件，这样可以篡改
   var err  error                   //源码可以看见，error的类型本质上是string
   if check_file_is_exist(f1.filename){   //看是否存在这个文件
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
```
