package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

type Hello struct {
	i int
	f float32
}

func main() {
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
	rv = reflect.ValueOf(&hello)
	if rv.CanSet() {
		fmt.Println("before set", rv, hello)
		rv.Elem().FieldByName("i").SetInt(5)
		fmt.Println("end    set", rv, hello)
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
