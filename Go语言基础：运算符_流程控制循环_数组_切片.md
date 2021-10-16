[toc]



# 运算符

## 算数运算符

| +    | 相加 |
| ---- | ---- |
| -    | 相减 |
| *    | 相乘 |
| /    | 相除 |
| %    | 求余 |

## 关系运算符

| ==   | 检查两个值是否相等，如果相等返回 True 否则返回 False。       |
| ---- | ------------------------------------------------------------ |
| !=   | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。   |
| >    | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。   |
| >=   | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 |
| <    | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。   |
| <=   | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 |

## 位运算符

| &    | 参与运算的两数各对应的二进位相与。 （两位均为1才为1）        |
| ---- | ------------------------------------------------------------ |
| \|   | 参与运算的两数各对应的二进位相或。 （两位有一个为1就为1）    |
| ^    | 参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。 （两位不一样则为1） |
| <<   | 左移n位就是乘以2的n次方。 “a<<b”是把a的各二进位全部左移b位，高位丢弃，低位补0。 |
| >>   | 右移n位就是除以2的n次方。 “a>>b”是把a的各二进位全部右移b位。 |

```go
package main

import "fmt"

// 位运算符
func main() {
   a:=5 //0101
   b:=6 //0110
   //位运算符使用的是二进制运算

   // 0101与0110按位与，得到0100=4
   fmt.Println(a&b)
   // 0101与0110按位或，得到0111=7
   fmt.Println(a|b)
   // 0101与0110按位异或，得到0011=3
   fmt.Println(a^b)
   // 0101向左偏移6位，得到0101000000=256+64
   fmt.Println(a<<b)
   // 如果偏移量超过了数据类型的范围，则多的位数丢弃，比如
   var c int8=64 
   fmt.Println(c<<b) //64为1000000 偏移6位为 1000000000000超过了8位的范围，所以高位舍弃
   //  0101与0110向右便宜6位，已经将101偏移出范围了，所以为0
   fmt.Println(a>>b)
}
```

## 逻辑运算符

| &&   | 逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False。 |
| ---- | ------------------------------------------------------------ |
| \|\| | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False。 |
| !    | 逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True。 |

## 赋值运算符

| =    | 简单的赋值运算符，将一个表达式的值赋给一个左值 |
| ---- | ---------------------------------------------- |
| +=   | 相加后再赋值                                   |
| -=   | 相减后再赋值                                   |
| *=   | 相乘后再赋值                                   |
| /=   | 相除后再赋值                                   |
| %=   | 求余后再赋值                                   |
| <<=  | 左移后赋值                                     |
| >>=  | 右移后赋值                                     |
| &=   | 按位与后赋值                                   |
| \|=  | 按位或后赋值                                   |
| ^=   | 按位异或后赋值                                 |







# 流程控制



## if语句

if语句的写法相比较于C语言，条件判断的时候没有括号。相比较于Python语言，在条件执行语句的时候具有花括号括起

```go
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

```

if条件判断还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断。这种写法会导致此变量的**作用域只在条件判断内部**，如下所示

```go
func main(){
	//特殊写法
	if your_age:=18;your_age<18{
		fmt.Println("you are too young!")
	}else if 18<=your_age && your_age<35{
		fmt.Println("it is the best")
	}else{
		fmt.Println("you are too old to do this job")
	}
	fmt.Println("your age is",your_age)
}

```

此时会报错：

```
# command-line-arguments
.\if02.go:15:28: undefined: your_age
```

## for循环

Go语言中只有for循环，没有while或者其他循环.

条件表达式返回`true`时循环体不停地进行循环，直到条件表达式返回`false`时自动退出循环。

一般形式,`此时i的作用域是for循环体`

```go
func main(){
	for i:=0 ;i<10;i++{
		fmt.Print(i," ")
	}
}
//0 1 2 3 4 5 6 7 8 9 
```

for循环的初始语句可以被忽略，但是初始语句后的分号必须要写，例如：

```go
func main() {
	i:=0
	for ;i<10;i++{
		fmt.Print(i)
	}
}
```

for循环的初始语句和结束语句都可以省略，逐渐向while循环转换，在书写时要防止死循环

```go
func main() {
	i:=0
	for i<10{
		fmt.Print(i)
		i++
	}
}
```

## for range 键值循环

Go语言中可以使用`for range`遍历数组、切片、字符串、map 及通道（channel）。即可以遍历类似于字典这种数据结构， 通过`for range`遍历的返回值有以下规律：

1. 数组、切片、字符串返回索引和值。
2. map返回键和值。
3. 通道（channel）只返回通道内的值。



使用range需要有两个变量来接收，如果使用一个变量接收的话，就只接收索引

```go
func main() {
	var l =[3]int{1,2,3}
	for a,b :=range l{
		fmt.Printf("%d:%d\n",a,b)
	}
}
//0:1
//1:2
//2:3
```

一个变量进行接收
```go
func main() {
   var l =[3]int{1,2,3}
   for b :=range l{
      fmt.Printf("%d\n",b)
   }
}
//0
//1
//2
```

## Swith 语句

使用`switch`语句可方便地对大量的值进行条件判断。

```go
func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}
```

Go语言规定每个`switch`只能有一个`default`分支。

一个分支可以有多个值，多个case值中间使用英文逗号分隔。

```go
func testSwitch2() {
	switch n := 7; n { //注意此时n的作用域
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}
```

分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：

```go
func switchDemo4() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}
```



`fallthrough`语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。（不常用）

```go
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
```



## goto跳转（不常用）

`goto`语句通过标签进行代码间的无条件跳转。`goto`语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用`goto`语句能简化一些代码的实现过程。 

```go
package main

import "fmt"

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}
func main() {
	gotoDemo2()
}
```

## break和continue

`break`语句可以结束`for`、`switch`和`select`的代码块。

`break`语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的`for`、`switch`和 `select`的代码块上。 举个例子：

```go
func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
```

`continue`语句可以结束当前循环，开始下一次的循环迭代过程，仅限在`for`循环内使用。

在 `continue`语句后添加标签时，表示开始标签对应的循环。例如：

```go
func continueDemo() {
forloop1:
   for i := 0; i < 5; i++ {
      // forloop2:
      for j := 0; j < 5; j++ {
         if i == 2 && j == 2 {
            continue forloop1
         }
         fmt.Printf("%v-%v\n", i, j)
      }
   }
}
```



# 数组

数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。即当时定义了容量为多大的数组，这个容量就不可再变化

## 数组定义

Go语言数组的定义和其他语言都不太一致，具体格式为：

```
var 数组名 [数组大小]数组内成员数据类型
```

比如

```go
var array [3]int
```

这就定义了名为array，大小为3，int类型的数组



## 数组初始化

### 直接/一般初始化

数组在定义的时候就可以直接初始化：

格式为

```
var 数组名 = [数组大小]数组内成员数据类型{数组值且用逗号分隔}
数组名 := [数组大小]数组内成员数据类型{数组值且用逗号分隔}
```

例如

```go
var array= [3]int{1,2,3}
fmt.Print(array)
//[1 2 3]
a := [3]int{10, 20, 30}
```



### 部分/默认值初始化

==如果不初始化数组或者部分初始化，那么没初始化的成员会具有默认值==

```go
func main() {
   // 数组定义
   var array[3]int
   var array_string[3]string
   var array_string_2=[4]string{"上海","北京"}
   var array_bool[3]bool
   var array_int=[3]int{1,2}
   fmt.Println(array)
   fmt.Println(array_string)
   fmt.Println("部分初始化",array_string_2)
   fmt.Println(array_bool)
   fmt.Println("部分初始化：",array_int)
}
//[0 0 0]
//[  ]
//部分初始化 [上海 北京  ]
//[false false false]
//部分初始化： [1 2 0]
```

### 不确定初始化

==当数组的大小不确定时，可以用[...]进行省略，系统会自动分配合适的空间给数组==

```go
package main

import "fmt"

func main() {
   var array[3]int
   var array_string_2=[...]string{"上海","北京"}
   var array_bool[3]bool
   var array_int=[...]int{1,2}
   fmt.Println(array)
   fmt.Println("不确定初始化",array_string_2)
   fmt.Println(array_bool)
   fmt.Println("不确定初始化：",array_int)
}
```

此外，不确定符号只能用于初始化，不能用于定义，如下的行为就是错误的

```go
var array_string[...]string
```



### 索引值初始化

也可以直接通过索引值给数组初始化：

```go
func main() {
   var array=[7]int{1:2,5:100,2:10}
   var array2=[...]int{1:2,20:10}
   fmt.Print(array," ")
   fmt.Printf("type of array2:%T \n",array)
   fmt.Print(array2," ")
   fmt.Printf("type of array2:%T",array2)
}
//[0 2 10 0 0 100 0] type of array2:[7]int
//[0 2 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 10] type of array2:[21]int
```

## 数组遍历

数组的遍历无异于要使用循环结构，下面介绍两种循环结构遍历数组

```go
package main

import "fmt"

func main() {
   var array=[...]int{1,2,3,4,5,6,7,8,9,10}
   //len函数可以获取数组长度
   fmt.Println("使用for循环遍历数组array")
   for i:=0;i<len(array);i++{
      fmt.Print(array[i]," ")
   }
   fmt.Println("\n使用for range循环遍历数组")
   for _,a:=range array{
      fmt.Print(a," ")
   }
}
//使用for循环遍历数组array
//1 2 3 4 5 6 7 8 9 10
//使用for range循环遍历数组
//1 2 3 4 5 6 7 8 9 10 
```

## 多维数组

以二维数组为例

### 二维数组定义和初始化

和其他语言的二维数组定义模板类似：

```go
var array=[3][2]int{{1,2},{3,4},{5,6}}
```

此外，定义的时候只有第一层可以使用[...]，例如

```go
var array2=[...][2]int{{1,2},{3,4},{5,6},{7,8}}
fmt.Println(array2) //[[1 2] [3 4] [5 6] [7 8]]
fmt.Println(len(array2)) //4
```

这种行为时不可取的

```go
var array2=[...][...]int{{1,2},{3,4},{5,6},{7,8}}
var array3=[4][...]int{{1,2},{3,4},{5,6},{7,8}}
```

### 二维数组遍历

两层循环嵌套，这里演示一下for和for range循环嵌套

```go
package main

import "fmt"

func main() {
   var array2=[...][2]int{{1,2},{3,4},{5,6},{7,8}}
   for i:=0;i<len(array2);i++{
      for _,a:=range array2[i]{
         fmt.Print(a,"\t")
      }
      fmt.Println()
   }
}
//1	2
//3	4
//5	6
//7	8	
```

## 值类型和引用类型

- **数组是值类型**，即当数组进行复制的时候，会额外开辟一个内存空间，这两个内存空间互不干扰，一方发生变化不会影响到另一方。
- 引用类型指的是指针，即浅复制，开辟一个内存空间不存放原来的数值，而是存放该目标的地址，一方发生变化，就会影响到另一方。

```go
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
```





# 切片

切片是一个难点

在python中，切片运用于序列，截取一段序列的值

比如上面的函数

```go
func change(a [3]int)  {
   a[2]=22
   fmt.Println("在函数体内部时",a)
}
```

这个函数必须时长度为3的数组，但是我们不好改变数组的长度，此时使用切片可以解决这个问题

## 切片的定义

切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含`地址`、`长度`和`容量`。切片一般用于快速地操作一块数据集合。

声明切片

```
var 数组名 []数据类型
```

如下所列

```go
func main() {
   // 声明切片类型
   var a []string              //声明一个字符串切片
   var b = []int{}             //声明一个整型切片并初始化
   var c = []bool{false, true} //声明一个布尔切片并初始化
   //var d = []bool{false, true} //声明一个布尔切片并初始化
   fmt.Println(a)              //[]
   fmt.Println(b)              //[]
   fmt.Println(c)              //[false true]
   fmt.Println(a == nil)       //true
   fmt.Println(b == nil)       //false
   fmt.Println(c == nil)       //false
   // fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}
```

这样感觉切片就是一个长度可变的数组

切片和数组的最大区别就是，切片时引用类型



## 切片的长度和容量

切片的长度就是当前切片的长度，使用len()函数进行求值。

切片的容量要联系到切片对象数组的指针，意义就是当前的数组的切片指针可以移动的范围大小，使用cap()函数进行求值。

在后面代码中会具体展现



## 切片表达式

### 切片简单表达式

切片就好比刀切一样，作用对象是数组

切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。 切片表达式中的`low`和`high`表示一个索引范围（左包含，右不包含）。即左开右闭区间。

```go
package main

import "fmt"

func main() {
   var array=[6]int{1,2,3,4,5,6}
   s:=array[:3]
   fmt.Printf("s:%v\t length:%v cap:%v",s,len(s),cap(s))
    fmt.Printf("\n%v",array[:3])
}
//s:[1 2 3]	 length:3 	cap:6
```



在这里可以看到，切片和python里面的切片基本一致。

此外，索引的上限要根据其容量来判断

```go
package main

import "fmt"

func main() {
   a := [5]int{1, 2, 3, 4, 5}
   s := a[1:3]  // s := a[low:high] 此时指针指向的是1号位置，即[2 3]
   fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
   s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s) 此时开始以指针所在的地方为首，在[2 3 4 5]中切片，就是[5]
   fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}
```

在观察切片的时候，一定要注意其指针到底在什么地方

### 切片的完整表达式

对于数组，指向数组的指针，或切片a(**注意不能是字符串**)支持完整切片表达式：

```go
a[low : high : max]
```

```go
package main

import "fmt"

func main() {
   var array=[8]int{1,2,3,4,5,6,7,8}
   fmt.Printf("array: %v\t array[:3:5]:%v\t len(array[:3:5]):%v\t cap(array[:3:4]):%v\t",array,array[:3:5],len(array[:3:5]),cap(array[:3:4]))

}
//array: [1 2 3 4 5 6 7 8]	 array[:3:5]:[1 2 3]	 len(array[:3:5]):3	 cap(array[:3:4]):4	
```

完整切片表达式需要满足的条件是`0 <= low <= high <= max <= cap(a)`，其他条件和简单切片表达式相同。

## make函数构造切片

我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的`make()`函数，格式如下：

```bash
make([]T, size, cap)
```

其中：

- T:切片的元素类型
- size:切片中元素的数量
- cap:切片的容量

```go
func main() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10
}
```

上面代码中`a`的内部存储空间已经分配了10个，但实际上只用了2个。 容量并不会影响当前元素的个数，所以`len(a)`返回2，`cap(a)`则返回该切片的容量。



## 切片的本质

切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

举个例子，现在有一个数组`a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}`，切片`s1 := a[:5]`，相应示意图如下。

![slice_01](https://www.liwenzhou.com/images/Go/slice/slice_01.png)

切片`s2 := a[3:6]`，相应示意图如下：![slice_02](https://www.liwenzhou.com/images/Go/slice/slice_02.png)





所以，切片的本质就是，len是切片的长度，cap就是当前指针指向的位置到底层数组的最后一个位置的长度



## 判断切片是否为空

要检查切片是否为空，请始终使用`len(s) == 0`来判断，而不应该使用`s == nil`来判断。

切片之间是不能比较的，我们不能使用`==`操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和`nil`比较。 一个`nil`值的切片**并没有底层数组**，一个`nil`值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是`nil`，例如下面的示例：

```go
var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
```

所以要判断一个切片是否是空的，要是用`len(s) == 0`来判断，不应该使用`s == nil`来判断。





## 切片赋值

由于是引用类型，是属于浅赋值

一方改变，另一方也会跟着改变

```go
package main

import "fmt"

func main() {
   array:=make([]int,3)
   s2:=array
   s2[2]=100
   fmt.Print(s2,"\n")
   fmt.Print(array)
}
//[0 0 100]
//[0 0 100]
```



## 切片遍历

切片遍历方法和数组一样

```go
func main() {
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}
```



## append方法为切片添加元素

Go语言的内建函数`append()`可以为切片动态添加元素。 可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（**后面加…**）。

```go
func main() {
   var array[]int
   array=append(array,1,2,3,4)
   fmt.Println(array)
   array2:=[]int{5,6,7,8}
   array=append(array,array2...) //后面必须为切片或其他值，不能是数组，如果是在后面要加...
   fmt.Println(array)
}
```



每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，**此时该切片指向的底层数组就会更换**。“扩容”操作往往发生在`append()`函数调用时，所以我们通常都需要用原变量接收append函数的返回值。

```go
package main

import "fmt"

func main() {
   var array[]int
   for i:=0;i<10;i++{
      array=append(array,i)
      fmt.Printf("%v len:%d cap:%d addr:%p \n",array,len(array),cap(array),array)
   }
}
//[0] len:1 cap:1 addr:0xc0000aa058
//[0 1] len:2 cap:2 addr:0xc0000aa0a0
//[0 1 2] len:3 cap:4 addr:0xc0000a80a0
//[0 1 2 3] len:4 cap:4 addr:0xc0000a80a0
//[0 1 2 3 4] len:5 cap:8 addr:0xc0000c20c0
//[0 1 2 3 4 5] len:6 cap:8 addr:0xc0000c20c0
//[0 1 2 3 4 5 6] len:7 cap:8 addr:0xc0000c20c0
//[0 1 2 3 4 5 6 7] len:8 cap:8 addr:0xc0000c20c0
//[0 1 2 3 4 5 6 7 8] len:9 cap:16 addr:0xc0000d4080
//[0 1 2 3 4 5 6 7 8 9] len:10 cap:16 addr:0xc0000d4080 
```

从上面的结果可以看出：

1. `append()`函数将元素追加到切片的最后并返回该切片。
2. 切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。





## 使用copy函数复制切片

切片是引用类型，寻常赋值进行复制切片传递的只是一个指针

使用copy函数可以做到开辟新的内存空间并进行值的复制

格式如下：

```bash
copy(destSlice, srcSlice []T)
```

其中：

- srcSlice: 数据来源切片
- destSlice: 目标切片

比如：

```go
package main

import "fmt"

func main() {
   var array=[]int{1,2,3,4,4,5,6,7,8}
   array2:=array //浅复制，指针
   array2[3]=100
   fmt.Println(array)
   fmt.Println(array2)

   array3:=make([]int,8,8)
   copy(array3,array) // 复制，开辟内存空间传递值
   array3[5]=1000
   fmt.Println(array3)
   fmt.Println(array)
}
//[1 2 3 100 4 5 6 7 8]
//[1 2 3 100 4 5 6 7 8]
//[1 2 3 100 4 1000 6 7]
//[1 2 3 100 4 5 6 7 8]
```





## 删除切片元素

go语言中没有删除切片的函数，但是可以利用数据结构的删除方法删除函数

比如

```go
package main

import "fmt"

func main() {
   var array=[]int{1,2,3,4,5,6,7}
   fmt.Println(array)
   array=append(array[:2],array[3:]...)
   fmt.Println(array)

}
//[1 2 3 4 5 6 7]
//[1 2 4 5 6 7]
```

总结一下就是：要从切片a中删除索引为`index`的元素，操作方法是`a = append(a[:index], a[index+1:]...)`





# 练习

下面的运行结果：

```go
package main

import "fmt"

func main() {
   var a = make([]string, 5, 10)
   fmt.Println(a) //[    ]
   for i := 0; i < 10; i++ {
      a = append(a, fmt.Sprintf("%v", i))
   }
   fmt.Println(a) //[     0 1 2 3 4 5 6 7 8 9]
}
```















































