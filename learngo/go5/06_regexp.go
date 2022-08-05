package main

import (
	"fmt"
	"regexp"
)

func findFloat(str string) (res []string) {
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("err")
	}
	ret := reg.FindAllStringSubmatch(str, -1)
	for _, x := range ret {
		res = append(res, x[0])
	}
	return
}
func main() {
	buf := "abc azc a7c aac 88 a9c tac"
	// 定义规则
	// reg := regexp.MustCompile(`a.c`)
	reg := regexp.MustCompile(`a\dc`)
	if reg == nil {
		fmt.Println("err")
		return
	}
	// 根据规则查找关键信息
	res := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println(res)

	//test1  匹配小数
	str := "3.14 123 asd 1.23 7. 5.6 sdfj 6.6 34.22"
	ret := findFloat(str)
	fmt.Println(ret)

	//test2 匹配网页

}
