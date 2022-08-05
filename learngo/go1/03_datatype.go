package main

import "fmt"

// type char byte

func main() {
	//变量定义   var [varname] [vartype]
	var a int
	fmt.Println("a = ", a)
	//变量初始化
	var b int = 10
	fmt.Println("b = ", b)

	// 赋值
	b = 20
	fmt.Println("b = ", b)

	//自动推导类型      :=  推导类型并赋值
	c := b // c 不需要提前定义  同一变量只能使用一次
	//c := 10 等同于 var c int =10
	fmt.Printf("c type is %T\n", c)

	var ch byte
	var str string
	var i int
	var f float32
	var f6 float64
	fmt.Println(ch, str, i, f, f6)

	type char byte //定义一个新的类型
	var cc char
	fmt.Printf("cc type is %T\n", cc)
}
