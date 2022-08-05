package calc

import "fmt"

func init() { //导包自动调用
	fmt.Println("import calc")
	fmt.Println("init")
}

func add(a, b int) int { //函数名小写 私有
	return a + b
}

func Add(a, b int) int { //函数名大写 公有
	// return a + b
	return add(a, b)
}
