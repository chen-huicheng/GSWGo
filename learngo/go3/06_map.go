package main

import "fmt"

func main() {
	// var m1 map[int]string
	dict := map[int]string{} //map[keytype]valuetype 基本类型 不能是切片
	dict[2] = "hello"
	dict[1] = "world"
	fmt.Println(dict)

	m2 := make(map[int]string)
	fmt.Println(m2)
	fmt.Println(len(m2))

	m3 := make(map[int]string, 5) //初始长度
	fmt.Println(m3)
	fmt.Println(len(m3))

	m4 := map[int]string{1: "hello", 2: "world"}
	fmt.Println(m4)

	//map 遍历
	for key, value := range m4 {
		fmt.Println(key, value)
	}

	// 判断 key 是否存在
	value, ok := m4[0]
	if ok == true {
		fmt.Println("ok ", value)
	} else {
		fmt.Println("key 不存在")
	}

	//map 删除
	fmt.Println(m4)
	delete(m4, 1)
	fmt.Println(m4)
	// map[1:hello 2:world]
	// map[2:world]

	delmap(m4, 2)
	fmt.Println(m4)

}

func delmap(m map[int]string, key int) { //map 传递的是引用
	delete(m, key)
}
