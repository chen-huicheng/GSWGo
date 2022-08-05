package main

import (
	"fmt"
	"io"
	"os"
)

type Humaner interface {
	sayhi()
}

type Personer interface {
	Humaner
	sing(lrc string)
}
type Student struct {
	name string
	id   int
}

func (this *Student) sayhi() {
	// this.name = "hello"
	fmt.Printf("student %+v\n", this)
}

type Teacher struct {
	name string
	id   int
}

func (this *Teacher) sing() {
	fmt.Printf("teacher sing %+v\n", this)
}

func WhoSayHi(i Humaner) {
	i.sayhi()
}
func main() {
	var i Humaner
	//只要实现了该的变量就可以该i赋值

	//cannot use s2 (variable of type Student) as Humaner value in assignment: missing method sayhi (sayhi has pointer receiver)
	// s2 := Student{"w", 5}  接口 i sayhi方法的receiver是 *Student
	// func (this *Student) sayhi() 是为 *Student 实现了sayhi()
	// 所以必须使用 i = &s2
	// 如果想使用 i=s2 需要实现 func (this Student) sayhi()
	// i = s2

	i = &Student{"c", 4}
	i.sayhi()
	// fmt.Printf("%+v\n", i)

	s := Student{"chen", 24}
	t := Teacher{"me", 24}
	WhoSayHi(&s)
	// WhoSayHi(&t)
	t.sing()

	fmt.Printf("\n\n\n\n")
	test()
}

func test() {
	op, err := os.OpenFile("go.go", os.O_RDWR, 0)
	defer op.Close()
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 10)
	for i := range buf {
		buf[i] = byte('a') + byte(i)
	}
	var iw io.Writer
	iw = op
	iw.Write(buf)

	rd_buf := make([]byte, 10)
	var ip io.Reader
	ip = iw.(io.Reader) //
	ip.Read(rd_buf)
	fmt.Println(string(rd_buf))

}
