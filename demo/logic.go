package main

import "fmt"

// 不使用 控制反转
// type A struct{}
// type B struct{}

// func (a *A) Controller() {
// 	b := new(B) //A 直接 new 来获得依赖对象 B 的引用
// 	b.Server()
// }
// func (b *B) Server() {
// 	fmt.Println("no IoC")
// }
// func main() {
// 	a := new(A)
// 	a.Controller()
// }

// 使用 控制反转
type A struct {
	b *B
}
type B struct{}

func (a *A) Controller() {
	a.b.Server()
}
func (b *B) Server() {
	fmt.Println("use IoC")
}
func initA() *A { // 依赖注入控制程序
	b := new(B)
	return &A{b}
}
func main() {
	a := initA()
	a.Controller()
}
