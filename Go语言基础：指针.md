[toc]



# 指针

区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。

任何程序数据载入内存后，在内存都有他们的地址，这就是指针。而为了保存一个数据在内存中的地址，我们就需要指针变量。

Go语言中的指针不能进行偏移和运算，因此Go语言中的指针操作非常简单，我们只需要记住两个符号：`&`（取地址）和`*`（根据地址取值）。

## 指针地址和指针类型

每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用`&`字符放在变量前面对变量进行“取地址”操作。 Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：`*int`、`*int64`、`*string`等。

取变量指针的语法如下：

```go
ptr := &v    // v的类型为T
```

其中：

- v:代表被取地址的变量，类型为`T`
- ptr:用于接收地址的变量，ptr的类型就为`*T`，称做T的指针类型。*代表指针。

举个例子：

```go
package main

import "fmt"

func main() {
	a:=10
	b:=10
	c:=b
	fmt.Printf("变量a的值为：%v 地址为：%p\n",a,&a)
	fmt.Printf("变量b的值为: %v 地址为：%p",b,&b)
//	%p占位符是取地址占位符
	fmt.Printf("变量c的值为: %v 地址为：%p",c,&c)
}

```

可见，虽然他们值相同，但是地址不同，每一个变量都有属于自己的地址

```
变量a的值为：10 地址为：0xc00000a0a8
变量b的值为: 10 地址为：0xc00000a0c0
变量c的值为: 10 地址为：0xc00000a0c8
```

即使存在赋值操作，也会开辟新的内存空间



## 指针赋值

下面例子表示，指针赋值b:=&a，于是b成为了一个指针

```go
package main

import "fmt"

func main() {
   a:=10
   b:=&a //b成了指针类型
   fmt.Printf("变量a的值为：%v 地址为：%p 类型为：%T\n",a,&a,a)
   fmt.Printf("变量b的值为: %v 地址为：%p 类型为：%T",b,&b,b)
   // %p占位符是取地址占位符
}
```

输出：

```ABAP
变量a的值为：10 地址为：0xc0000aa058 类型为：int
变量b的值为: 0xc0000aa058 地址为：0xc0000ce018 类型为：*int
```

可以发现，b的值为a的地址，而b是在一个新开辟的内存空间里面

![取变量地址图示](https://www.liwenzhou.com/images/Go/pointer/ptr.png)

所以它是不是一个指针，首先必须是带*的指针类型，然后值是地址



## 指针取值

对指针使用*就可以对其取值了：

```go
package main

import "fmt"

func main() {
   a:=10
   b:=&a //b成了指针类型
   fmt.Printf("变量a的值为：%v 地址为：%p 类型为：%T\n",a,&a,a)
   fmt.Printf("变量b的值为: %v 地址为：%p 类型为：%T\n",b,&b,b)
   fmt.Printf("b是a的指针，从b那里可以取得a的值，值为*b=%v",*b)
   // %p占位符是取地址占位符
}
```

结果：

```nsis
变量a的值为：10 地址为：0xc00000a0a8 类型为：int
变量b的值为: 0xc00000a0a8 地址为：0xc000006028 类型为：*int
b是a的指针，从b那里可以取得a的值，值为*b=10
```



取地址操作符`&`和取值操作符`*`是一对互补操作符，`&`取出地址，`*`根据地址取出地址指向的值。

变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：

- 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
- 指针变量的值是指针地址。
- 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。



## 指针传值

指针传递值的时候是地址的指向，修改指针所指向的值是到其原地址修改

这个机制和C语言一样

如下所示：

```go
package main

import "fmt"

func ptr06_a(x int){
   x=20
   fmt.Println("在函数ptr06_a中x的地址",&x)
}
func ptr06_b(x *int){
   *x=20
   fmt.Println("在函数ptr06_b中x的地址",&x)
   fmt.Println("在函数ptr06_b中*x的地址",&*x)
}
func main() {
   x:=100
   fmt.Println("全局变量x的地址：",&x)
   ptr06_a(x)
   fmt.Println(x)
   ptr06_b(&x) //*int是指针类型，指针的值是地址，所以需要传递地址
   fmt.Println(x)
   fmt.Println("最后x的地址：",&x)
}
```

```
全局变量x的地址： 0xc00012a058
在函数ptr06_a中x的地址 0xc00012a080
100
在函数ptr06_b中x的地址 0xc00014e020
在函数ptr06_b中*x的地址 0xc00012a058
20
最后x的地址： 0xc00012a058
```

函数ptr06_a中x的地址和全局变量x的地址不一样，所以任凭修改x都没有用

函数ptr06_b中x的地址是零一个指针，指针的值是0xc00012a058，*x才是指针指向的值，所以修改\*x,就是修改地址为0xc00012a058的值

所以甚至可以写出，指针的**伪代码**：

```go
func find_ptr(x){
   if x is ptr{
      if value(x)!=nil{
         for _,address in range [all address]{
             if value(x)==address{
               return value(x)
            }
         }
         panic("Can't find address")
      }
      else{
         panic("you have not value of address")
      }
   }
   
}
```

## new和make

我们先来看一个例子：

```go
func main() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
```

上面的指针a没有指向任何东西，盲目得给a所指向的东西赋值

就像你想敲键盘上α键，但是键盘上又没有，你心里不就开始报错了？

在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存。

## new

new是一个内置的函数，它的函数签名如下：

```go
func new(Type) *Type
```

其中，

- Type表示类型，new函数只接受一个参数，这个参数是一个类型
- *Type表示类型指针，new函数返回一个指向该类型内存地址的指针。

new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。举个例子：

```go
func main() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}	
```

本节开始的示例代码中`var a *int`只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。应该按照如下方式使用内置的new函数对a进行初始化之后就可以正常对其赋值了：

new是开辟了一个新的内存空间供指针初始化，指针的空间早已开辟了

```go
package main

import "fmt"

func main() {
	var a *int
	a=new(int)
	*a=10
	fmt.Printf("a的地址为 %p \na的值为 %v\n*a的地址为 %p\n*a的值为 %v",&a,a,&*a,*a)
}
```

```go
a的地址为 0xc000006028 
a的值为 0xc00000a0a8
*a的地址为 0xc00000a0a8
*a的值为 10
```

## make

make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。make函数的函数签名如下：

```go
func make(t Type, size ...IntegerType) Type
```

make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。这个我们在上一章中都有说明，关于channel我们会在后续的章节详细说明。

本节开始的示例中`var b map[string]int`只是声明变量b是一个map类型的变量，需要像下面的示例代码一样使用make函数进行初始化操作之后，才能对其进行键值对赋值：

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var l map[string]int32
	fmt.Printf("l的地址是：%p\n",&l)
	fmt.Printf("l的大小是：%v\n",unsafe.Sizeof(l))
    //初始化
	l=make(map[string]int32,10)
	l["aaa"]=122
	fmt.Printf("l[aaa]的大小是：%v",unsafe.Sizeof(l["aaa"]))
}
//l的地址是：0xc0000ce018
//l的大小是：8
//l[aaa]的大小是：4
```

在未初始化的时候，l的大小已经有了，是8，但是内存一旦分配好就没法变动，它的内存是8，这明显是不够的，所以后面使用make的时候实际上是开辟了一个新的内存空间，来对l进行一个索引，所以l是属于引用类型的。



## new与make的区别

1. 二者都是用来做内存分配的。
2. make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3. 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。



