package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	for i := 0; i <= 100; i++ { //++运算符只能使用在变量后边
		sum += i
	}
	fmt.Println("sum = ", sum)

	k := 0
	println("time.second", time.Second)
	println("time.minute", time.Minute)
	println("time.hour", time.Hour)
	println("time.now.day", time.Now().Day())
	for { //死循环
		time.Sleep(time.Second)
		if k > 5 {
			break
		}
		k++
		fmt.Println("k = ", k)
	}

	str := "123456789"
	for i, c := range str { // i 是元素位置 下标  c 是元素本身
		fmt.Printf("idx : %d, char : %c\n", i, c)
	}

	//下面两种 等同
	for i := range str { // i 是元素位置  第二个抛弃
		fmt.Printf("idx : %d, char : %c\n", i, str[i])
	}
	// 不使用不赋值
	// for i, _ := range str { // i 是元素位置  第二个抛弃
	// 	fmt.Printf("idx : %d, char : %c\n", i, str[i])
	// }
}
