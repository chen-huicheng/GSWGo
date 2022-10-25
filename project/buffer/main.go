package main

import (
	"bytes"
	"fmt"
	"strings"
)

func byteBufferDemo() {
	//创建
	var b bytes.Buffer        //直接定义一个 Buffer 变量，而不用初始化
	b.Write([]byte("Hello ")) // 可以直接使用
	fmt.Printf("%s\n", b.String())

	b1 := new(bytes.Buffer)    //直接使用 new 初始化，可以直接使用
	b1.Write([]byte("Buffer")) // 可以直接使用
	fmt.Printf("%s\n", b1.String())
	// 其它两种定义方式
	// func NewBuffer(buf []byte) *Buffer
	// func NewBufferString(s string) *Buffer
	buf := make([]byte, 10)
	// b3 := bytes.NewBuffer([]byte("hello byte buffer"))
	b3 := bytes.NewBufferString("hello byte buffer") // 与上一行等效
	n, err := b3.Read(buf)
	if err != nil {
		fmt.Printf("Read error;err=%s\n", err)
		return
	}
	fmt.Println("read buffer", n, buf[:n], b3.String())

	// 可以对 Buffer 进行读写

}
func stringsBuilderDemo() {
	var b strings.Builder
	for i := 1; i <= 3; i++ {
		fmt.Fprintf(&b, "%d...", i)
		b.WriteString(fmt.Sprintf("###%d ", i))
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
}
func main() {
	byteBufferDemo()
	stringsBuilderDemo()
}
