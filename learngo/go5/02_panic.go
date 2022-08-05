package main

import "fmt"

func test() {
	fmt.Println("this is test")
	panic("this is a panic") //类似 c++  assert()
}
func test1(x int) {
	var a [10]int
	a[x] = 11

}
func main() {
	// test()
	test1(15)
	//panic
	//"runtime error: index out of range [15] with length 10"
}
