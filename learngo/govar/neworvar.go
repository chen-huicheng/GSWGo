package main

import "fmt"

var p *int
var 变量 int

func main() {
	p = heap()
	fmt.Printf("&p:%p\n", p)
	stack()

	变量 = 10
	fmt.Println(变量)
}

func stack() {
	a := 1
	b := new(int)
	fmt.Printf("&a:%p --- b:%p\n", &a, b)
}

func heap() *int {
	a := 1
	b := 2
	fmt.Printf("&a:%p --- &b:%p\n", &a, &b)
	return &b
}

func ass() {
	var a int = 10
	var b float32 = 2.2
	a = int(b)
	b = float32(a)

}
