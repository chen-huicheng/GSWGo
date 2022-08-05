# 内存管理
## 1.垃圾回收机制
[参考1](https://zhuanlan.zhihu.com/p/297177002) []()
### why
> 不进行垃圾回收可能导致内存泄漏。
> 语言层面的垃圾回收，使的程序开发更便捷，提高开发效率。

垃圾回收器主要包括三个目标：

> **无内存泄漏**：垃圾回收器最基本的目标就是减少防止程序员未及时释放导致的内存泄漏，垃圾回收器会识别并清理内存中的垃圾
> **自动回收无用内存**：垃圾回收器作为独立的子任务，不需要程序员显式调用即可自动清理内存垃圾
> **内存整理**：如果只是简单回收无用内存，那么堆上的内存空间会存在较多碎片而无法满足分配较大对象的需求，因此垃圾回收器需要重整内存空间，提高内存利用率

### how
> 垃圾回收主要分为两种算法“**引用计数法**”和“**追踪回收**”

+ **引用计数**
> 根据每个对象的引用计数器是否为0来判断该对象是否为未引用的垃圾对象
> 

+ **追踪回收**
> 先判断哪些对象存活，然后将其余的所有对象作为垃圾进行回收

### what


## 2.反射机制
反射第一定律：反射可以将interface类型变量转换成反射对象
反射第二定律：反射可以将反射对象还原成interface对象
反射第三定律：反射对象可修改，value值必须是可设置的
[code](./../go09/01_reflect.go)
```go
package main
import (
	"fmt"
	"reflect"
)
type iface interface{}
type Hello struct {
	i int
	f float32
}
func main() {
	var i iface
	i = Hello{1, 2.0}
	fmt.Println("type:", reflect.TypeOf(i))
	fmt.Println("value:", reflect.ValueOf(i))
}
// 输出
// type: main.Hello
// value: {1 2}
```

# 3.闭包
### 基本概念
闭包是可以包含自由(未绑定到特定对象)变量的代码块，这些变量不在这个代码块内或者 任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块(由于自由变量包含 在代码块中，所以这些自由变量以及它们引用的对象没有被释放)为自由变量提供绑定的计算环 境(作用域)。
> 理解：闭包是包含环境变量的匿名函数。
```go
package main
import ( "fmt")
func main() {
	var j int = 5
	a := func()(func()) { 
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		} 
	}()
	a()
	j *= 2 
	a()
}
//输出
//i, j: 10, 5
//i, j: 10, 10
```
### 闭包的价值 
闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，就是说这些函数可以存储到 变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。
### Go语言中的闭包 
Go语言中的闭包同样也会引用到函数外的变量。闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在，

# 4.包管理
### go package
> go通过包来管理程序
> 引入一个包 使用 import pkg_name
> 通常一个文件夹下的所有文件属于同一个包，**包名**为文件夹名称。
> 导入时使用路径导入。
> 例如：
> $GOPARH/src/project/dirA/dirB/pkg_dir/a.go b.go c.go
> 使用
> import "project/dirA/dirB/pkg_dir" //导入
> pkg_dir.func_name() //调用函数

### go module 发布管理
> 如果想发布go 模块 在github上。
> **首先** 初始化 go.mod 中 
> module github.com/chen-huicheng/GSWGo   发布网站/yourname/项目名称
> **然后导入** 当别人引入你的包时，通过  go get 发布网站/yourname/项目名称 引入 并通过import 
> 
> 通常go mod init 使用$GOPARH/src 到当前目录的路径初始化 module 即本模块的名字。当导入本包的文件时
