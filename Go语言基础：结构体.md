# 结构体

Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。

Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。

## 类型别名和自定义类型

### 自定义类型

在Go语言中有一些基本的数据类型，如`string`、`整型`、`浮点型`、`布尔`等数据类型， Go语言中可以使用`type`关键字来定义自定义类型。

自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。例如：

```go
//将MyInt定义为int类型
type MyInt int
```

通过`type`关键字的定义，`MyInt`就是一种新的类型，它具有`int`的特性。

### 类型别名

类型别名是`Go1.9`版本添加的新功能。

类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。

```go
type TypeAlias = Type
```

我们之前见过的`rune`和`byte`就是类型别名，他们的定义如下：

```go
type byte = uint8
type rune = int32
```

### 类型定义和类型别名的区别

类型别名与类型定义表面上看只有一个等号的差异，我们通过下面的这段代码来理解它们之间的区别。

```go
package main

import "fmt"

//类型定义和类型别名的区别
func main() {
   type MyInt int
   type cname_int=int

   var a MyInt
   var b cname_int
   fmt.Printf("Type of MyInt: %T\n",a)
   fmt.Printf("Type of cname_int: %T",b)

}
```

效果：

```
Type of MyInt: main.MyInt
Type of cname_int: int
```

可以看到，别名并不能改变数据类型，而使用自定义类型就可以改变数据类型。

结果显示a的类型是`main.MyInt`，表示main包下定义的`MyInt`类型。b的类型是`int`。`cname_int`类型只会在代码中存在，编译完成时并不会有`cname_int`类型。

## 结构体

Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了，Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称`struct`。 也就是我们可以通过`struct`来定义自己的类型了。

**Go语言中通过`struct`来实现面向对象。**

## 结构体的定义

使用`type`和`struct`关键字来定义结构体，具体代码格式如下：

```go
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
```

其中：

- 类型名：标识自定义结构体的名称，在同一个包内不能重复。
- 字段名：表示结构体字段名。结构体中的字段名必须唯一。
- 字段类型：表示结构体字段的具体类型。

举个例子，我们定义一个`Person`（人）结构体，代码如下：

```go
type person struct {
	name string
	city string
	age  int8
}
```

同样类型的字段也可以写在一行，

```go
type person1 struct {
	name, city string
	age        int8
}
```

这样我们就拥有了一个`person`的自定义类型，它有`name`、`city`、`age`三个字段，分别表示姓名、城市和年龄。这样我们使用这个`person`结构体就能够很方便的在程序中表示和存储人信息了。

语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型

## 结构体实例化

**只有当结构体实例化时，才会真正地分配内存。**也就是必须实例化后才能使用结构体的字段。

结构体本身也是一种类型，我们可以像声明内置类型一样使用`var`关键字声明结构体类型。

相当于面向对象中的创建一个对象

```go
var 结构体实例 结构体类型
```

### 基本实例化

举个例子：

```go
package main

import "fmt"

type person struct {
   name string
   age int
   sex string
}

func main() {
   var p1 person
   p1.name="张三"
   p1.age=23
   fmt.Println(p1.name,p1.sex,p1.age)
   fmt.Printf("p1=%#v",p1)
}
```

效果：

```
张三  23
p1=main.person{name:"张三", age:23, sex:""}
```

可见，初始化的方法是`对象.属性=值`,未被初始化的属性存在并且为空

### 匿名结构体

格式：

```
var 对象名 struct{属性名 属性类型,属性名 属性类型……}
```

在定义一些临时数据结构等场景下还可以使用匿名结构体。

```go
package main

import "fmt"

func main() {
   var p1 struct{name string;age int}
   p1.name="李四"
   p1.age=33
   fmt.Printf("%#v\n",p1)
   fmt.Println(p1.name,p1.age)
}
```

效果：

```
struct { name string; age int }{name:"李四", age:33}
李四 33
```

可以发现，此时结构体并没有名字，所以这是匿名结构体

### 创建指针类型结构体

我们还可以通过使用`new`关键字对结构体进行实例化，得到的是结构体的地址。 格式如下：

```go
var p2 = new(person)
fmt.Printf("%T\n", p2)     //*main.person
fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
```

从打印的结果中我们可以看出`p2`是一个结构体指针。

需要注意的是在Go语言中支持对结构体指针直接使用`.`来访问结构体的成员。

```go
package main

import "fmt"

type person04 struct{
	name string
	age int
	sex string
}
func main() {
	var p1=new(person04)
	p1.sex="女"
	p1.name="小红"
	p1.age=20
	fmt.Printf("p1: %#v\n",p1)
	fmt.Printf("value of p1:%v\n",*p1)
	fmt.Printf("value of p1:%p\n",p1)
	fmt.Printf("address of p1:%p\n",&p1)
	fmt.Printf("address of *p1:%p\n",&*p1)

	var p2 person04
	p2.name="王五"
	p2.sex="男"
	fmt.Printf("value of p2:%v\n",p2)
	fmt.Printf("address of p2:%p",&p2)
}

```

效果：

```go
p1: &main.person04{name:"小红", age:20, sex:"女"}
value of p1:{小红 20 女}
value of p1:0xc000076480
address of p1:0xc000006028
address of *p1:0xc000076480
value of p2:{王五 0 男}
address of p2:0xc000076510
```

可以看到，如果是用new开辟的空间，就把此对象当作指针来使用，如上p1的值是&{小红 20 女}，就是说{小红 20 女}的地址0xc000076480，，如果是正常声明的空间，就当作值来使用

需要注意的是在Go语言中支持对结构体指针直接使用`.`来访问结构体的成员。

p2.name实际上是(*p2).name，下一个例子会说明这个问题

### 取结构体的地址实例化

使用`&`对结构体进行取地址操作相当于对该结构体类型进行了一次`new`实例化操作。

```go
package main

import "fmt"

type person05 struct{
	name string
	age int
	sex string
}
func main() {
	p2:=&person05{}
	p2.sex="nv"
	p2.name="Emil"
	fmt.Printf("type of p2 : %T\n",p2)
	fmt.Printf("value of p2: %v\n",p2)
	fmt.Printf("value of p2: %v\n",*p2)
	(*p2).age=22
	(*p2).name="Have"
	fmt.Printf("value of p2: %v\n",*p2)
}
```

## 结构体初始化

没有初始化的结构体，其成员变量都是对应其类型的零值。

```go
package main

import "fmt"

type person06 struct{
   name string
   age int
   sex string
}
func main() {
   var p1 person06
   fmt.Printf("value of p1: %#v",p1)
}
```

```
value of p1: main.person06{name:"", age:0, sex:""}
```



### 使用键值对初始化

- 使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。

- 也可以对结构体指针进行键值对初始化

- 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。

- 初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值

  使用这种格式初始化时，需要注意：

  1. 必须初始化结构体的**所有字段**。
  2. 初始值的填充顺序必须与字段在结构体中的声明**顺序一致**。
  3. 该方式**不能和键值初始化方式混用**。



例子

```go
package main

import "fmt"

type person07 struct{
	name string
	age int
	sex string
}
func main() {
	//对结构体初始化
	p1:=person07{
		name:"Emak",
		age:22,
		sex:"male",
	}
	fmt.Printf("value of p1:%v\n",p1)
	//对结构体指针初始化
	p2:=&person07{
		name:"Alice",
		sex:"female",
	}
	fmt.Printf("value of p2:%v",*p2)
	//省略键的结构体初始化赋值
	p3:=&person07{
		"Pbied",
		22,
		"male",
	}
	fmt.Printf("value of p3:%v",*p3)
}

```

效果：

```
value of p1:{Emak 22 male}
value of p2:{Alice 0 female}
value of p3:{Pbied 22 male}
```



## 结构体内存布局

### 占用连续的内存。

```go
package main

import "fmt"

type person08 struct{
   a int8
   b int8
   c int8
   d int8
}
func main() {
   n:=person08{
      1,2,3,4,
   }
   fmt.Printf("n.a %p\n", &n.a)
   fmt.Printf("n.b %p\n", &n.b)
   fmt.Printf("n.c %p\n", &n.c)
   fmt.Printf("n.d %p\n", &n.d)
}
```

```
n.a 0xc0000aa058
n.b 0xc0000aa059
n.c 0xc0000aa05a
n.d 0xc0000aa05b
```

可见，占用的字符是连续的，int8代表的是8个字节，就是1kb，即一块内存空间

### 空结构体

空结构体是不占用空间的。

```go
var v struct{}
fmt.Println(unsafe.Sizeof(v))  // 0
```





## 面试题

```go
type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
```



## 构造函数

Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个`person`的构造函数。 因为`struct`是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。



```go
package main

import "fmt"

type person10 struct {
   name string
   age  int
}

func newPerson10(name string,age int) *person10{
   return &person10{
      name:name,
      age:age,
   }
}
func main() {
   p10:=newPerson10("张三",23)
   fmt.Printf("value of p10: %#v",*p10)
}
//value of p10: main.person10{name:"张三", age:23}
```

## 方法和接收者

Go语言中的`方法（Method）`是一种作用于特定类型变量的函数。这种特定类型变量叫做`接收者（Receiver）`。接收者的概念就类似于其他语言中的`this`或者 `self`。

方法的定义格式如下：

```go
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

其中，

- 接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是`self`、`this`之类的命名。例如，`Person`类型的接收者变量应该命名为 `p`，`Connector`类型的接收者变量应该命名为`c`等。
- 接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
- 方法名、参数列表、返回参数：具体格式与函数定义相同。

例子

```go
package main

import "fmt"

type person11 struct{
   name string
   age int
}
//构造函数
func newPerson11(name string,age int) *person11{
   return &person11{
      name:name,
      age:age,
   }
}
//方法
func (p person11) tell11(){
   fmt.Printf("%s is %v",p.name,p.age)
}

func main() {
   var p11=newPerson11("小红",22)
   p11.tell11()

}
```

这里的：

```go
func (p person11) tell11(){
   fmt.Printf("%s is %v",p.name,p.age)
}
```

就是定义的方法

方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。



### 指针和值类型的接收者

指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。这种方式就十分接近于其他语言中面向对象中的`this`或者`self`。 例如我们为`Person`添加一个`SetAge`方法，来修改实例变量的年龄。

而不使用指针进行修改，无法达到修改的预期效果，即当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。

```go
package main

import "fmt"

type person11 struct{
	name string
	age int
}
//构造函数
func newPerson11(name string,age int) *person11{
	return &person11{
		name:name,
		age:age,
	}
}
//方法

//指针类型的接收者
func (p *person11) change_age(new_age int){
	p.age=new_age //相当于this.age=new_age
}
func (p person11) tell11(){
	fmt.Printf("%s is %v",p.name,p.age)
}

//值类型的接收者
func (p person11) change_name(new_name string){
	p.name=new_name
	//fmt.Println(p.name)
}
func main() {
	var p11=newPerson11("小红",22)
	p11.change_age(18)
	p11.change_name("小王")
	p11.tell11()

}
//小红 is 18
```

### 什么时候应该使用指针类型接收者

1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

## 任意类型添加方法

在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。 举个例子，我们基于内置的`int`类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

```go
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
```

**注意事项：** 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。

## 结构体的匿名字段

结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。

**注意：**这里匿名字段的说法并不代表没有字段名，而是**默认会采用类型名作为字段名**，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个，如果非得出现相同类型的（比如说有多个字符串属性），可以**声明别名或者自定义类型**

```go
package main

import "fmt"

type Int8 int8
type Int64=int64
type Person13 struct{
   int
   int64
   int8
   Int8
   Int64
}
func main() {
   p1:=Person13{
      1,2,3,4,5,
   }
   fmt.Printf("value of p1: %#v",p1)

}
//value of p1: main.Person13{int:1, int64:2, int8:3, Int8:4, Int64:5}
```

## 嵌套结构体

结构体可以嵌套新的结构体：

比如下面的结构体中，如果次要信息和主要信息放在同一个结构体中会非常难看，

可以将次要信息嵌套在主要信息当中

```go
package main

import "fmt"

type person14 struct {
   name string
   age int
   infor infor14
}
type infor14 struct {
   address string
   telephone int64
}

func main() {

   p:=person14{
      name:"Tom",
      age:33,
      infor:infor14{
         address:"Fire Street",
         telephone: 1228847389,
      },
   }
   fmt.Printf("The information of p is :%#v",p)
}
//The information of p is :main.person14{name:"Tom", age:33, infor:main.infor14{address:"Fire Street", telephone:1228847389}}
```

### 嵌套匿名字段

上面person结构体中嵌套的infor结构体也可以采用匿名字段的方式，例如：

```go
package main

import "fmt"

type person15 struct {
	name string
	age int
	infor15
}
type infor15 struct {
	address string
	telephone int64
}

func main() {

	var p person15
	p.infor15.telephone=34567890 // 两层赋值
	p.name="甄嬛"
	p.address="紫阳小街" //匿名字段名可以省略
	fmt.Printf("The information of p is :%#v",p)
}

//The information of p is :main.person15{name:"甄嬛", age:0, infor15:main.infor15{address:"紫阳小街", telephone:34567890}}
```

当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。

### 嵌套结构体的字段名冲突

嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名。

```go
package main
//Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	Address
	Email
}

func main() {
	var user3 User
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
}
```

## 结构体的“继承”

Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。

```go
package main

import "fmt"
// 继承
type Animal17 struct {
   name string
}
func (a *Animal17) move(){
   fmt.Printf("%s 会动\n",a.name)
}
type Chicken17 struct {
   food string
   *Animal17 //通过匿名结构体的嵌套实现了继承
}
// 定义方法
func (a *Chicken17) eat(){ //嵌套的是指针，这样的话可以节省大量空间
   fmt.Printf("%s 会吃 %s\n",a.name,a.food)
}
func main() {
   a:=&Chicken17{
      food :"粮食",
      Animal17:&Animal17{
         name:"鸭子",
      },
   }
   a.move()
   a.eat()
}
```

## 结构体字段的可见性

结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

（没想到吧，还能这么定义）

```go
package main

import (
   "encoding/json"
   "fmt"
)

type student19 struct {
   ID int
   Sex string
   name string //私有不能被json包访问
}


func main() {
   s1:=student19{
      ID:1,
      name:"aaa",
      Sex:"man",
   }
   data, err :=json.Marshal(s1)
   if err !=nil{
      fmt.Println("json序列化出现错误")
      return
   }
   fmt.Printf("json :%s\n",data)

}
//json :{"ID":1,"Sex":"man"}
```

## 结构体与JSON序列化

JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。易于人阅读和编写。同时也易于机器解析和生成。JSON键值对是用来保存JS对象的一种方式，键/值对组合中的键名写在前面并用双引号`""`包裹，使用冒号`:`分隔，然后紧接着值；多个键值之间使用英文`,`分隔。

```go
package main

import (
   "encoding/json"
   "fmt"
)

type Student struct {
   ID int
   Gender string
   Name string
}

type Class struct {
   Title string
   Students []*Student //定义结构体切片，因为一个班上会有很多学生，使用指针可以减少内存消耗
}


func main() {
   c:=&Class{
      Title: "计算机科学与技术1班",
      Students: make([]*Student,0,200), //结构体切片也需要初始化
   }
   for i:=0 ;i<10 ;i++{
      stu:=&Student{
         Name: fmt.Sprintf("student:%02d",i),
         Gender:"man",
         ID:i,
      }
      c.Students=append(c.Students,stu)
   }
   //JSON 序列化：即生成JSON格式字符串
   data,error:=json.Marshal(c)
   if error!=nil{
      fmt.Println("json序列化出现错误")
      return
   }
   fmt.Printf("jion:%s\n",data)

   //JSON反序列化
   str:=data
   c1:=&Class{} //创建一个空的Class对象，用于接收值
   error=json.Unmarshal([]byte(str),c1)
   if error!=nil{
      fmt.Println("json反序列化出现错误")
      return
   }
   fmt.Printf("%#v\n",c1)


}
```

## 结构体标签（Tag）

`Tag`是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 `Tag`在结构体字段的后方定义，由一对**反引号**包裹起来，具体的格式如下：

```bash
`key1:"value1" key2:"value2"`
```

结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔。

**注意事项：** 为结构体编写`Tag`时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在key和value之间添加空格。

例如我们为`Student`结构体的每个字段定义json序列化时使用的Tag：

```go
package main

import (
   "encoding/json"
   "fmt"
)

type student19 struct {
   ID int `json:"id"` //在json中ID字段就会变成id
   Sex string
   name string //私有不能被json包访问
}


func main() {
   s1:=student19{
      ID:1,
      name:"aaa",
      Sex:"man",
   }
   data, err :=json.Marshal(s1)
   if err !=nil{
      fmt.Println("json序列化出现错误")
      return
   }
   fmt.Printf("json :%s\n",data)

}
//json :{"id":1,"Sex":"man"}
```

## 结构体和方法补充知识点

因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意。我们来看下面的例子：

```go
type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	p.dreams = dreams
}

func main() {
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams)  //  [吃饭 不睡觉 打豆豆]
}
```

可以看见，由于data是切片类型，修改了data的值，一位置那个对象的dreams的值被修改了

下面使用了copy函数进行了值复制，可以解决这个问题

```go
package main


import "fmt"

type Person21 struct {
   name   string
   age    int8
   dreams []string
}

func (p *Person21) SetDreams(dreams []string) {
   p.dreams = make([]string, len(dreams))
   copy(p.dreams, dreams) //值复制
}

func main() {
   p1 := Person21{name: "小王子", age: 18}
   data := []string{"吃饭", "睡觉", "打豆豆"}
   p1.SetDreams(data)

   // 你真的想要修改 p1.dreams 吗？
   data[1] = "不睡觉"
   fmt.Println(p1.dreams) // [吃饭 不睡觉 打豆豆]

}
```
