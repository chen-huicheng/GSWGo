package gotest3_test

import "github.com/chen-huicheng/GSWGo/learngo/gotest3"

// 检测单行输出
func ExampleSayHello() {
	gotest3.SayHello()
	// OutPut: Hello World
}

// 检测多行输出
func ExampleSayGoodbye() {
	gotest3.SayGoodbye()
	// OutPut:
	// Hello,
	// goodbye
}

// 检测乱序输出
func ExamplePrintNames() {
	gotest3.PrintNames()
	// Unordered output:
	// Jim
	// Bob
	// Tom
	// Sue
}
