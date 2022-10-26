package app

import "fmt"

// import "fmt"

// type A struct {
// 	b *B
// }
// type B struct{}

// func (a *A) Controller() {
// 	a.b.Server()
// }
// func (b *B) Server() {
// 	fmt.Println("use IoC")
// }
// func NewA(b *B) *A {
// 	return &A{b}
// }
// func NewB() *B {
// 	return &B{}
// }

type A struct {
	is IServer // A 依赖接口 IServer
}
type IServer interface {
	Server()
}
type B struct{}

func (a *A) Controller() {
	a.is.Server()
}
func (b *B) Server() { // B 实现了 IServer
	fmt.Println("use IoC")
}
