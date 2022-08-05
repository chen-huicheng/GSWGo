package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

type long int

func (tmp long) Add02(other long) long { //给这个类挂载一个函数
	return tmp + other
}

func test1() {
	res := Add(1, 2)
	fmt.Println(res)

	var a long = 5
	var b long = 4
	ret := a.Add02(b) //调用方式
	fmt.Println(ret)
}

// 继承
type Person struct {
	name string
	sex  byte
	age  int
}

func (this Person) PrintInfo() {
	fmt.Printf("info %+v\n", this)
}

func (this Person) setInfo1(name string, sex byte, age int) { //值传递，无法修改对象内部值
	this.name = name
	this.sex = sex
	this.age = age
}
func (this *Person) setInfo2(name string, sex byte, age int) {
	this.name = name
	this.sex = sex
	this.age = age
}
func main() {
	test1()
	p := Person{"chen", 'm', 24}
	p.PrintInfo()

	var p2 Person
	p2.setInfo1("h", 'f', 34)
	fmt.Println("p2")
	p2.PrintInfo()

	// (&p2).setInfo2("h", 'f', 34)
	p2.setInfo2("h2", 'f', 35) //等价于上
	p2.PrintInfo()

	// info {name:chen sex:109 age:24}
	// info {name: sex:0 age:0}
	// info {name:h sex:102 age:34}

	//方法值
	pfunc := p2.PrintInfo //use of value  将方法作为一个值使用
	pfunc()               //然后调用

	//方法表达式
	f := (*Person).PrintInfo //
	f(&p2)
}
