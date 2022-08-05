package main

import (
	"fmt"
	"math/rand"
)

func sort(arr [10]int) { //数组做函数参数 是值传递
	n := len(arr)
	flag := false
	for i := n - 1; i > 0; i-- {
		flag = true
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				flag = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if flag {
			break
		}
	}
	fmt.Println("arr = ", arr)
}
func sort1(arr *[10]int) { //数组做函数参数 是值传递
	n := len(*arr)
	flag := false
	for i := n - 1; i > 0; i-- {
		flag = true
		for j := 0; j < i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				flag = false
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
		if flag {
			break
		}
	}
	fmt.Println("arr = ", *arr)
}
func main() {
	//设置种子
	//种子一样 则每次程序运行结果都一样
	rand.Seed(666)

	//产生随机数
	fmt.Println("rand = ", rand.Int()) //rand =  4343637058903381868
	fmt.Println("randn = ", rand.Intn(100))

	arr := [10]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	sort(arr)
	fmt.Println("arr = ", arr) // 还是原来的arr 没有排序

	sort1(&arr)
	fmt.Println("arr = ", arr) //使用指针传递
}
