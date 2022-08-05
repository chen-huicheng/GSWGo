package main

import (
	"fmt"
	"unsafe"
)

type pointer struct {
	pi   *int
	pb   *byte
	a    int
	b    byte
	pii  *int16
	c    int16
	d    int16
	e    byte
	pi32 *int32
}

func test() {
	p := pointer{}
	fmt.Printf("sizeof ß%v\n", unsafe.Sizeof(p))
	fmt.Println(&p.pi)
	fmt.Println(&p.pb)
	fmt.Println(&p.a)
	fmt.Println(&p.b)
	fmt.Println(&p.pii)
	fmt.Println(&p.c)
	fmt.Println(&p.d)
	fmt.Println(&p.e)
	fmt.Println(&p.pi32)

}

func main() {
	test()
	a := 10
	fmt.Printf("a = %d\n", a)
	fmt.Printf("&a = %v\n", &a)

	//保存变量地址
	var p *int //默认 = <nil>
	fmt.Println("p = ", p)
	//不能对 空指针操作
	// *p = 666 //invalid memory address or nil pointer dereference

	//p指针的使用
	p = &a
	fmt.Printf("*p = %d\n", *p)
	fmt.Printf("p = %v\n", p)

	*p = 666
	fmt.Printf("*p = %d\n", *p)
	fmt.Printf("p = %v\n", p)

	//new 使用   申请一块内存
	var pi *int
	pi = new(int)
	*pi = 77
	fmt.Printf("*pi = %d\n", *pi)
	fmt.Printf("pi = %v\n", pi)

	c, d := 1, 2
	fmt.Printf("main:c=%d,d=%d\n", c, d)
	swap(&c, &d)
	fmt.Printf("swap:c=%d,d=%d\n", c, d)
}
func swap(a, b *int) {
	*a, *b = *b, *a
}
