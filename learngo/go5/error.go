package main

import (
	"errors"
	"fmt"
	"io"
)

type MyError struct {
	Msg string
}

func (e MyError) Error() string {
	return "MyError:" + e.Msg
}

var Me = MyError{"me"}

func test() {
	e := io.EOF
	fmt.Println(e == io.EOF)          //true
	fmt.Println(errors.Is(e, io.EOF)) //true
	fmt.Println(e)
	e = fmt.Errorf("context: %w", e)
	fmt.Println(e == io.EOF)          //false
	fmt.Println(errors.Is(e, io.EOF)) //true
	fmt.Println(e)
}

func test1() {
	e := Me
	fmt.Println(errors.As(e, &Me)) //true
	fmt.Println(e)
	ne := fmt.Errorf("context: %w", e)
	fmt.Println(errors.As(ne, &Me)) //true
	fmt.Println(ne)
}
func main() {
	test()
	test1()
}
