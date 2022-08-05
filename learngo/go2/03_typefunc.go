package main

import "fmt"

func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}

type FuncType func(int, int) int

func Calc(a, b int, ftest FuncType) (res int) {
	res = ftest(a, b)
	return
}

func main() {
	res := Add(1, 2)
	fmt.Println(res)
	res = Sub(3, 4)
	fmt.Println(res)

	// var ftest FuncType
	ftest := Add //函数类型也可推理得到

	fmt.Printf("ftest : %T\n", ftest)
	res = ftest(1, 2)
	fmt.Println(res)

	ftest = Sub
	res = ftest(3, 4)
	fmt.Println(res)

	res = Calc(3, 4, Add)
	fmt.Println(res)

}
