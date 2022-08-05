package main

import "fmt"

func main() {
	const alen int = 10
	var id [alen]int //[数字] 数值必须是常量

	// c := 10
	// var id2 [c]int //array length c (variable of type int) must be constant

	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d] = %d\n", i, id[i])
	}

	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	c := [5]int{1, 2, 3} //未初始化部分为 0
	fmt.Println("c = ", c)

	d := [5]int{2: 10, 4: 3}
	fmt.Println("d = ", d)
	// a =  [1 2 3 4 5]
	// b =  [1 2 3 4 5]
	// c =  [1 2 3 0 0]
	// d =  [0 0 10 0 3]

	var arr [3][4]int
	k := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			k++
			arr[i][j] = k
			fmt.Printf("arr[%d][%d]=%d\n", i, j, arr[i][j])
		}
	}
	arr1 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {8, 9, 10, 11}}
	fmt.Println("arr1 = ", arr1)

	//数组比较 仅支持 ==  !=   数组类型一样，每个元素一样

	//同类型可赋值

	arr1 = arr
	fmt.Println("arr1 = ", arr1)
}
