package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

type Hello struct {
	N int
	f float32
}

func test1(mod interface{}, v int) interface{} {
	modType := reflect.TypeOf(mod).Elem()
	fmt.Println(modType)
	dst := reflect.New(modType).Interface()
	fmt.Println(dst)
	return dst
}

func test2() {
	type t struct {
		N int
		s string
	}
	n := t{42, "hello"}
	fmt.Println("before", n)
	rv := reflect.ValueOf(&n).Elem()
	rv.FieldByName("N").SetInt(7)
	fmt.Println("after", n)
}
func main() {
	test2()
	a := test1(&Hello{}, 2)
	fmt.Println(a)
	var i interface{}
	hello := Hello{1, 2.0}
	i = hello
	// 反射第一定律：反射可以将interface类型变量转换成反射对象
	fmt.Println("type:", reflect.TypeOf(i))
	fmt.Println("value:", reflect.ValueOf(i))
	// 反射第二定律：反射可以将反射对象还原成interface对象
	rv := reflect.ValueOf(i)
	ii := rv.Interface()
	if hello1, ok := ii.(Hello); ok {
		fmt.Println(hello1)
	}
	if i == ii {
		fmt.Println("interface{hello} == reflect.ValueOf(hello).Interface()")
	} else {
		fmt.Println("!=")
	}

	// 反射第三定律：反射对象可修改，value值必须是可设置的  即指针类型
	if rv.CanSet() {
		fmt.Println("hello can set")
	}

	rvx := reflect.ValueOf(&hello).Elem()
	// reflect.ValueOf(&hello).Elem().FieldByName("i").SetInt(7)
	if rvx.CanSet() {
		fmt.Println("before set", rvx, hello)
		rvx.FieldByName("N").SetInt(5)
		fmt.Println("end    set", rvx, hello)
	}
	test()
}

func test() {
	fmt.Println("\ntest func")
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	var w io.Writer
	// w.Write([]byte("hello"))  //panic: runtime error: invalid memory address or nil pointer dereference

	w = os.Stdout
	w.Write([]byte("hello\n"))

	w = new(bytes.Buffer)
	w.Write([]byte("world\n"))

	fmt.Println(w)

	var a interface{}
	a = 3
	fmt.Println(a)
}
