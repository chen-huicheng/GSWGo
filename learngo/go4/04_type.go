package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3) //使用一个空接口 来存储变量  可以存储任何类型的变量
	i[0] = 1
	i[1] = "str"
	i[2] = Student{"c", 33}
	// fmt.Printf("type %T\n", i[2])
	for idx, data := range i {
		if v, ok := data.(int); ok { // 类型断言  var.(vartype)
			fmt.Printf("i[%d] is int; v = %d\n", idx, v)
			// fmt.Printf("type %T\n", v)
		} else if v, ok := data.(string); ok {
			fmt.Printf("i[%d] is string; v = %s\n", idx, v)
		} else if v, ok := data.(Student); ok {
			fmt.Printf("i[%d] is Student; v = %+v\n", idx, v)
		}
	}
	for idx, data := range i {
		switch v := data.(type) { //
		case int:
			fmt.Printf("i[%d] is int	; v = %d\n", idx, v)
		case string:
			fmt.Printf("i[%d] is string	; v = %s\n", idx, v)
		case Student:
			fmt.Printf("i[%d] is Student ; v = %+v\n", idx, v)
		}

	}
}
