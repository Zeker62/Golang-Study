[toc]





# map

map类似与python中的字典，由键值对构成

Go语言中的map是引用类型，必须初始化才能使用。

## map的定义

Go语言中 `map`的定义语法如下：

```go
map[键]值
```

map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

```go
make(map[KeyType]ValueType, [cap])
```

其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

## map使用

有两种使用方法，一是开辟内存空间再赋值，二是直接赋值

```go
name:=make(map[string]int)
name["张三"]=23
name["李四"]=25
fmt.Println(name)
fmt.Println(name["李四"])
fmt.Printf("type of map:%T\n",name)
 //直接赋值
address:=map[string]string{
   "Tom":"break",
   "Alice":"continue",
}
fmt.Println(address)
```

输出结果

```
map[张三:23 李四:25]
25
type of map:map[string]int
map[Alice:continue Tom:break]
```



不可以没有开辟空间再根据索引赋值，例如

```go
age:=map[string]string
age["Cab"]="aaa"
fmt.Println(age)
```

这是错误的



## 判断某个键是否存在

```go
value, is := map[key]
```

使用此格式的话，如果key存在，is变量会返回true，如果key不存在，那么is变量会返回false，value变量会返回0

```go
package main

import "fmt"

func isOK(ok bool,v int){
   if ok{
      fmt.Println(v,ok)
   }else{
      fmt.Println(v,ok)
      fmt.Println("没有此键")
   }
}
func main() {
   name:=make(map[string]int,8)
   name["aaa"]=123
   name["bbb"]=456
   v,ok:=name["ccc"]
   isOK(ok,v)
   v,ok=name["aaa"]
   isOK(ok,v)
}
//0 false
//没有此键
//123 true
```



## map的遍历

遍历非常简单，和数组的遍历类似

```go
func main() {
   maps:=make(map[string]int,10)
   maps["aaa"]=123
   maps["bbb"]=345
   maps["ccc"]=678
   for a,b:=range maps{
      fmt.Println("maps[",a,"]=",b)
   }
}
```

当只想遍历key的时候，循环体可以改为：

```go
for a:=range maps{
   fmt.Println("maps[",a,"]")
}
```

此外，需要注意的是：==map遍历的顺序和赋值的顺序无关==

即在第一个代码中，每次运行输出的结果不同。

## 使用delete函数删除键值对

使用`delete()`内建函数从map中删除一组键值对，`delete()`函数的格式如下：

```go
delete(map, key)
```

- map:表示要删除键值对的map
- key:表示要删除的键值对的键

上代码：

```go
package main

import "fmt"

func display(maps map[string]int){
   for a,b:=range maps{
      fmt.Println("maps[",a,"]",b)
   }
}
func main() {
   maps:=make(map[string]int,10)
   maps["bbb"]=345
   maps["ccc"]=678
   maps["aaa"]=123
   maps["ddd"]=0

   fmt.Println("删除前")
   display(maps)
   fmt.Println("删除ccc")
   delete(maps,"ccc")
   display(maps)
}
//maps[ bbb ] 345
//maps[ ccc ] 678
//maps[ aaa ] 123
//maps[ ddd ] 0
//删除ccc
//maps[ aaa ] 123
//maps[ ddd ] 0
//maps[ bbb ] 345
```

这个例子不仅可以看出delete函数的作用，还可以验证之前所说：map的遍历顺序和它的添加顺序无关



## 特定的顺序遍历map

这种方法就是先将map中的键存在一个切片当中，然后对切片进行排序，再通过索引的方式输出value，以达到遍历map的目的

```go
package main

import (
   "fmt"
   "math/rand"
   "sort"
   "time"
)
func main() {
   // 初始化随机种子
   rand.Seed(time.Now().UnixNano())

   var maps=make(map[string]int,200)
   for i:=0;i<50;i++{
      key:=fmt.Sprintf("student:%02d",i)
      value:=rand.Intn(200)
      maps[key]=value
   }
   // 将key存入到keys切片当中去
   keys:=make([]string,0,200)
   for key:=range maps{
      keys=append(keys,key)
   }
   fmt.Println(keys)

   sort.Strings(keys)
   for _,key:=range keys{
      fmt.Println(key,"=",maps[key])
   }

}
```



## 元素是map类型的切片

定义切片的时候，数据类型是map

```go
package main

import "fmt"

func main() {
	// 切片里面有map
	slice :=make([]map[string]string,3)
	for a,b :=range slice {
		fmt.Printf("key:%d value:%v\n",a,b)
	}
	fmt.Println("初始化……")
	slice[0] =make(map[string]string,4)
	slice[0]["name"]="aaa"
	slice[0]["age"]="123"
	for index,value := range slice {
		fmt.Printf("key:%d value:%v\n",index,value)
	}
    fmt.Println(slice)
}

```

输出结果：

```
key:0 value:map[]
key:1 value:map[]
key:2 value:map[]
初始化……
key:0 value:map[age:123 name:aaa]
key:1 value:map[]
key:2 value:map[]
[map[age:123 name:aaa] map[] map[]]
```



## 元素类型是切片的map

如下所示，可以在map中添加类型为切片的值

```go
package main

import "fmt"

func main() {
   maps:=make(map[string][]string,5)
   fmt.Println(maps)
   fmt.Println("初始化……")
   key:="City"
   value,ok :=maps[key]
   if !ok{
      value=make([]string,0,6)
   }
   value=append(value,"Beijing","Shanghai")
   maps[key]=value
   fmt.Println(maps)

}
//map[]
//初始化……
//map[City:[Beijing Shanghai]]
```



## 练习

1. 写一个程序，统计一个字符串中每个字母出现的次数。

   ```go
   package main
   
   import "fmt"
   
   //写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
   func main() {
      s:="how do you do"
      array:=[]rune(s)
      maps:=make(map[string]int,10)
      for _,c:=range array{
         value,ok :=maps[string(c)]
         if ok{
            value+=1
            maps[string(c)]=value
         }else{
            maps[string(c)]=1
         }
      }
      for keys,values:=range maps{
         if keys!=" "{
            fmt.Printf("%v=%d ",keys,values)
         }
   
      }
   }
   ```

   

2. 观察下面代码，写出最终的打印结果。

```go
func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
```





# 函数

## 函数定义

Go语言中定义函数使用`func`关键字，具体格式如下：

```go
func 函数名(参数)(返回值){
    函数体
}
```

其中：

- 函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
- 参数：参数由参数变量和参数变量的类型组成，多个参数之间使用`,`分隔。
- 返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用`()`包裹，并用`,`分隔。
- 函数体：实现指定功能的代码块。


```go
func sum(a int,b int)int{
   sum:=a+b
   return sum
}
```



## 函数的调用

定义了函数之后，我们可以通过`函数名()`的方式调用函数。 例如我们调用上面定义的两个函数，代码如下：

```go
func sum(a int,b int)int{
	fmt.Println("求和函数")
	sum:=a+b
	return sum
}

func main() {
	fmt.Println(sum(3,5))
}
```

注意，调用有返回值的函数时，可以不接收其返回值。

## 参数类型的简写

**函数的参数中如果相邻变量的类型相同，则可以省略类型**，例如：

```go
func intSum(x, y int) int {
	return x + y
}
```

上面的代码中，`intSum`函数有两个参数，这两个参数的类型均为`int`，因此可以省略`x`的类型，因为`y`后面有类型说明，`x`参数也是该类型。

## 可变参数

可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加`...`来标识。

注意：可变参数通常要作为函数的最后一个参数。

```go
package main

import "fmt"

func sum002(x int,y...int)int{
    //注意，此时y是一个切片
	sum:=x
	for _,v :=range y{
		sum+=v
	}
	return sum
}
func main() {
	fmt.Println(sum002(1,2,3,4,5))
}

```





## 返回值

Go语言中通过`return`关键字向外输出返回值。

### 多返回值

Go语言中函数支持多返回值，函数如果有多个返回值时必须用`()`将所有返回值包裹起来。

```go
func ca001(x int,y int)(int,int){
   return x+y,x-y
}
func main() {
   fmt.Println(ca001(5,7)) //12 -2
}
```



### 返回值命名

函数定义时可以给返回值命名，并在函数体中具体直接使用这些变量，最后通过`return`关键字返回。

注意如下代码，此时已经定义了sum和sub，下面的符号用赋值符号

```go
func ca002(x int,y int)(sum ,sub int){
   sum=x+y
   sub=x-y
   return
}
func main() {
   fmt.Println(ca001(5,7)) //12 -2
}
```



### 返回值补充

当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，可以用作不显示返回一个长度为0的切片。

```go
func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	...
}
```







# 变量作用域

### 全局变量

全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在任何函数中可以访问到全局变量。

```go
package main

import "fmt"
var a int=200

func ca003(x int,y int)int{
   return a
}
func main() {
   fmt.Println(ca003(5,7)) //200
}
```

### 局部变量

全局变量不能通过函数来改变，其改变的范围仅限于其函数体

例如，函数内定义的变量无法在该函数外使用

main函数中无法使用testLocalVar函数中定义的变量x：

```go
func testLocalVar() {
	//定义一个函数局部变量x,仅在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

func main() {
	testLocalVar()
	fmt.Println(x) // 报错，此时无法使用变量x
}
```



如果局部变量和全局变量重名，优先访问局部变量。

下面这种情况局部变量和全局变量都用到了as，当函数体内发现自己定义变量时，发现自己和全局变量重名，会重新开辟一个内存空间放入值，并且as指针将指向这个值直到跳出函数体

```go
package main

import "fmt"
var as int=200

func ca004()int{
	as:=10
	return as
}
func main() {
	fmt.Println(as) //200
	fmt.Println(ca004()) //10
	fmt.Println(as) //200
}
```

如果变量在循环体，判断体中，其作用域也是其局部作用

```go
package main

import "fmt"

func main() {
	a:=[]int{1,2,3,4,5,6,7}
	for i:=0;i<len(a);i++{
		b:=i
		if b<10{
			c:=10
			fmt.Println(b,c)
		}
	}
	//fmt.Println(i)
	//fmt.Println(b)
	//fmt.Println(c)
    //以上三条报错
}
```

如上，i和b都不能在循环体外输出，因为这两个是在循环体内定义的。c也不能在判断体外输出。



# 函数类型与变量

## 定义函数类型

我们可以使用`type`关键字来定义一个函数类型，具体格式如下：

```go
type calculation func(int, int) int
```

上面语句定义了一个`calculation`类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。

简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。

```go
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
```

add和sub都能赋值给calculation类型的变量。

```go
var c calculation
c = add
```

## 函数类型变量使用

我们可以声明函数类型的变量并且为该变量赋值：

例如

```go
package main

import "fmt"

type calculate func(int,int)int
//定义这种形态的函数都是calculate类型

func sum(a int,b int)int{
   s:=a+b
   return s
}
func sub(x int,y int)int{
   s:=x-y
   return s
}


func main() {
   var c calculate
// 定义一个变量，类型是calculate样式的函数
   c=sum
// 可以将函数sum赋值给c
   fmt.Printf("type of c is %T\n",c)
   fmt.Println(c(10,20))

   // 没有经过calculate直接赋值
   f:=sub
   fmt.Printf("type of f is %T\n",f)
   fmt.Println(f(20,10))
}

```

结果

```
type of c is main.calculate
30
type of f is func(int, int) int
10
```

# 高级用法

### 函数作为参数

```go
package main

import "fmt"

func sum009(x int,y int)int{
   return x+y
}



func summary009(x int,y int,function func( int, int)int)int{
   return function(x,y)
}


func main() {
   result:=summary009(3,5,sum009)
   fmt.Println(result)
}

//8
```



### 函数作为返回值



```go
package main

import "fmt"

func sum010(x int,y int)(func (int,int)int){
   s:=x+y
   if s>100 {
      return sum010(s, y)
   }
   return nil
}


func main() {
   fmt.Println(sum010(5,3))
}
```





# 匿名函数

函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数的定义格式如下：

```go
func(参数)(返回值){
    函数体
}
```

匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

```go
func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
```

匿名函数多用于实现回调函数和闭包。

再举个例子

```go
package main

import "fmt"

func main() {
   sum:=func(x,y int){
      fmt.Println(x+y)
   }
   sum(10,20)
   fmt.Println(func(x,y int)int{
      return x+y+y
   }(10,20))

}
```

这是有返回值的写法

匿名函数的写法就是如下：

```
变量名:=func(变量 类型……)返回值{
	函数体
}
变量名(参数……)
------------------------
func(变量 类型)返回值{
	函数体
}(参数……)
```

























