package main

import "fmt"

/* 函数定义模版
func funcName(v1 valType1, v2 valType2, v3, v4 valType3)(r1 valType4, r2 valtype5){
}
*/
// 无参数无返回值函数
func MyFunc1() { // 在main前后都可执行
	// do
}

// 单参数函数
func MyFunc2(v int) {
	// do(v)
}

// 多参数函数
func MyFunc3(v1 int, v2 int) { // 等价 func MyFunc3(v1, v2 int) {
	// do(v1, v2)
}

//不定参数
func MyFunc4(args ...int) { //不定长参数只能在最后一个
	fmt.Println("len ", len(args))
	for i, item := range args {
		fmt.Printf("idx : %d, char : %d\n", i, item)
	}
}

// 有返回值 返回值类型写在 函数参数列表后 花括号前
func MyFunc5() int {
	return 666
}

//给返回值 起一个名字  常用写法
func MyFunc6() (ret int) {
	ret = 666
	return
}

// 返回多个返回值
// 等价于 func MyFunc7() (r1 int, r2 int, r3 int)
// func MyFunc7() (r1, r2, r3 int)
func MyFunc7() (int, int, int) {
	return 1, 2, 3
}
func MyFunc8() (v1 int, v2 int, v3 int) {
	v1, v2, v3 = 1, 2, 3
	return
}
func main() {
	MyFunc1()
	MyFunc2(2)
	MyFunc3(1, 2)
	MyFunc4(2, 3, 4, 5, 6)
	_ = MyFunc5()
	MyFunc6()
	MyFunc7()
	MyFunc8()

}
