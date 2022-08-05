package main

import "fmt"

func MyFunc() { //无参数无返回值函数
	fmt.Println("hello MyFunc")
}

func main() {
	MyFunc()
	MyFunc1()

	a := 10
	MyPrint(a)

	b := 5
	Add1(a, b)

	MyPrint(a, b)

	c := myreturn()
	d := myreturn2()
	fmt.Printf("c : %d, d : %d\n", c, d)

	a, b, c = myreturn3()
	fmt.Printf("a : %d, b : %d, c = %d\n", a, b, c)

	a, b, c = myreturn4()
	fmt.Printf("a : %d, b : %d, c = %d\n", a, b, c)

}

func MyFunc1() { //无参数无返回值函数   在main前后都可执行
	fmt.Println("hello MyFunc")
}

func MyPrintint(a int) {
	fmt.Println(a)
}

func Add1(a int, b int) { // 等价 func Add(a , b int) {
	fmt.Println("sum is ", a+b)
}

//不定参数
func MyPrint(args ...int) { //不定长参数只能在最后一个
	fmt.Println("len ", len(args))
	for i, item := range args {
		fmt.Printf("idx : %d, char : %d\n", i, item)
	}
}

// 有返回值 返回值类型写在 函数参数列表后 花括号前
func myreturn() int {
	return 666
}

//给返回值 起一个名字  常用写法
func myreturn2() (ret int) {
	ret = 666
	return
}

// 返回多个返回值
func myreturn3() (int, int, int) {
	return 1, 2, 3
}

func myreturn4() (a int, b int, c int) {
	a, b, c = 11, 22, 33
	return
}

func myreturn5() (a, b, c int) {
	a, b, c = 11, 22, 33
	return
}
