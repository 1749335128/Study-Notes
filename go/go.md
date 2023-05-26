## 

`go env -w GO111MODULE=auto`go语言环境配置找不到go.mod文件，开启感知模式

## 基本关键字与标识符

Go 代码中会使用到的 25 个关键字或保留字,Go 代码中的关键字保持的这么少，是为了简化在编译过程第一步中的代码解析,关键字不能用作标志符。

| 关键字      | 或保留字        |        |           |        |
| -------- | ----------- | ------ | --------- | ------ |
| break    | default     | func   | interface | select |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

Go 语言还有 36 个预定义标识符，其中包含了基本类型的名称和一些基本的内置函数。

| 标志符       |            |        |       |         |         |
| --------- | ---------- | ------ | ----- | ------- | ------- |
| append    | bool       | byte   | cap   | close   | complex |
| complex64 | complex128 | uint16 | copy  | false   | float32 |
| float64   | imag       | int    | int8  | int16   | int32   |
| uint32    | int64      | iota   | len   | make    | new     |
| nil       | panic      | uint64 | print | println | real    |
| recover   | string     | true   | uint  | uint8   | uintptr |

## 包的概念，导入和可见性

每个程序都由包（通常简称为 pkg）的概念组成，可以使用自身的包或者从其它包中导入内容。一个包可以由许多以 `.go` 为扩展名的源文件组成，文件名和包名一般来说都是不相同的。

**标准库**:在 Go 的安装文件里包含了一些可以直接使用的包，即标准库,go的标准库包含了大量的包

**如果对一个包进行更改或重新编译，所有引用了这个包的客户端程序都必须全部重新编译。**

包的导入方式

```go
import "fmt"
import "os"
import "fmt"; import "os"
import (
   "fmt"
   "os"
)
import ("fmt"; "os")//使用 gofmt 后将会被强制换行
```

如果包名不是以.或/开头，以./开头路径go会去全局文件进行查找，反之会到相对目录下查找，以/开头会在系统的绝对路径中查找。

**可见性规则**

当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头,这种形式的标志符可以被外部的包所使用，但必须导入这个包；当标志符是小写字母开头，对外部包是不可见的，但是整个包的内部是可见和使用的。

go语言要求导入的包必须使用，否则会引发构建错误。

## Go程序基本结构

**函数**

大括号必须放在函数名同行，返回类型写在方法后。

func Sum(a, b int) int { return a + b }

**类型**

类型可以是基本类型，如：int、float、bool、string；结构化的（复合的），如：struct、array、slice、map、channel；只描述类型的行为的，如：interface。

**常量**

常量使用关键字 `const` 定义，用于存储不会改变的数据

`const a int =10`

`const a = 10`

`const beef, two, c = "eat", 2, "veg"`

`const ( Monday, Tuesday, Wednesday = 1, 2, 3 Thursday, Friday, Saturday = 4, 5, 6 )`

```go
const ( 
a = iota 
b = iota 
c = iota )
```

**变量**

`a = 15 b = false`

`var a int = 15 var i = 5`

`var a`无法自动推`断`

`HOME = os.Getenv("HOME")`go语言在运行时自动推断类型

局部变量` a :=10`

**init函数**

变量可以在 init 函数中初始化。这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。

**布尔值**

```go
var aVar=10
var b bool = aVar == 10
fmt.Printf("%t",b)
```

**结构体**

```go
type identifier struct {
    field1 type1
    field2 type2
}
func main(){
    var i identifier
    p = new(identifier)
    fmt.Print(p.filed1)
}
    type struct1 struct{
    a int
    b float32
    c string
}
   ms := &struct1{10, 15.5, "Chris"}//此时ms 为struct1类型的指针
    ms := struct1{10,15.5,"chris"}//结构体变量
```

**复数**

```go
complex64 (32 位实数和虚数)
complex128 (64 位实数和虚数)
var c complex128=5 + 3i//使用
fmt.Printf("%v",c)
如果 re 和 im 的类型均为 float32，那么类型为 complex64 的复数 c 可以通过
以下方式来获得：c = complex(re, im)
real(c) //获得实数部分
imag(c) //获得虚数部分
```

**二元运算符**

1. 按位与 `&`

2. 按位或 `|`

3. 按位异或 `^`

**随机数**

```go
import ( "fmt" 
"math/rand" 
)
a := rand.Int()
b := rand.Intn(8)
timens := int64(time.Now().Nanosecond())
```

**类型别名**

`type TZ int`

**字符类型**

```go
var ch byte = 65  
var ch byte = '\x41'
```

**字符串类型**

在循环中使用字符串拼接，尽量不要使用+号

`strings.Join()`

`bytes.Buffer`

字符串的函数

如果字符串是非ASCII，建议使用

`strings.IndexRune(s string, r rune) int`

**多返回值函数**

```go
//多返回值函数,使用_下划线可以忽略函数的返回值
func getNumber(s string) (int, bool) {
   defer strings.Contains(s, "1")//本函数最后执行，return执行完之后再执行
    var a int
    var b error
    a, b = strconv.Atoi(s)
    var c bool = b == nil
    return a, c
}
```

**变长参数函数**

要求最后一个形参是...type的形式，参数的个数可以是0个获多个type类型变量

```go
func Greeting(prefix string, who ...string)
Greeting("hello:", "Joe", "Anna", "Eileen"
如果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}，
这样就可以接受任何类型的参数
func typecheck(values … interface{}) {

}
```

**defier语句和追踪**

键字 defer 允许我们推迟到函数返回之前,函数的最后才执行，一般用于关闭释放的资源，类似于Java中的finally用法

```go
func f() {
    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
    }
}//输出 4 3 2 1 0

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
    trace("a")
    defer untrace("a")
    fmt.Println("in a")
}

func b() {
    trace("b")
    defer untrace("b")
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
//用defer实现函数追踪
```

**内置函数**

1. close函数用于管道通信，关闭通道

2. len返回数据的长度

3. cap返回某个类型的最大存储量

4. new、make用于分配内存，new用于值类型和用户自定义类型(含自定义结构)，make用于内置引用类型(切片，管道，map)

**switch语句**

```go
//在go语言当中不适用break，也只会执行当前分支
switch {
    case condition1:
        ...
    case condition2:
        ...
    default:
        ...
}
```

**for语句**

```go
for 初始化语句;条件语句;修饰语句{

}
for i := 0; i < 5; i++ {
        fmt.Printf("This is the %d iteration\n", i)
}
```

```go
for 位置,对象 range 容器对象{

}
for pos, char := range str {
        fmt.Printf("Character on position %d is: %c \n", pos, char)
}
```

**标签与goto语句**

for、switch 或 select 语句都可以配合标签（label）形式的标识符使用，即某一行第一个以冒号（`:`）结尾的单词（gofmt 会将后续代码自动移至下一行）。

```go
//标签和 goto 语句之间不能出现定义新变量的语句，否则会导致编译失败
func main() {
    i:=0
    HERE:
        print(i)
        i++
        if i==5 {
            return
        }
        goto HERE
}//使用标签和 goto 语句是不被鼓励的
```

## gin web框架

下载并安装

```sh
 go get -u github.com/gin-gonic/gin
```

将gin引入到代码环境中

```go
import "github.com/gin-gonic/gin"
引入常量包(可选)
import "net/http"
```

##### 打包部署

windows平台

```
set CGO_ENABLED=0 //禁用CGO
set GOOS=linux //目标平台为linux
set GOARCH=amd64 //目标处理器架构是amd64
go build -o name //编译可执行文件到当前目录 （-o：自定义文件名）
```
