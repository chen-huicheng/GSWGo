package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// func Contains(s,substr string) bool //判断子串是否在s中
	fmt.Println(strings.Contains("seafood", "foo")) //true

	// func Join(a []string,sep string) string
	s := []string{"foo", "bar", "ba"}
	fmt.Println(strings.Join(s, ", ")) //foo, bar, ba

	// func Index(s,sep string) int //返回下标 不存在返回 -1
	fmt.Println(strings.Index("chicken", "ken")) //4
	fmt.Println(strings.Index("chicken", "kk"))  //-1

	// func Repeat(s string, count int) string // 返回拼接结果
	fmt.Println(strings.Repeat("hello ", 5)) //hello hello hello hello hello

	// func Replace(s, old, new string, n int) string // n 代表替换次数 <0 全部替换
	fmt.Println(strings.Replace("oink oink oink oink", "in", "of", 2))  //oofk oofk oink oink
	fmt.Println(strings.Replace("oink oink oink oink", "in", "of", -1)) //oofk oofk oofk oofk

	// func Split(s, sep string) []string
	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) //["a" "b" "c"]

	//func Trim(s, cutset string) string
	fmt.Printf("%q\n", strings.Trim(" ! ! ! hello !!! ", "! ")) //hello

	// func Fields(s string) []string //
	fmt.Printf("%q\n", strings.Fields("  foo bar hel   ")) //["foo" "bar" "hel"]

	// 字符串转换
	//Append 将其他类型转换为字符串 并添加到字符串
	//Format 将其他类型转换为字符串
	//Parse  将字符串转换为其他类型

	str := make([]byte, 0, 100)
	// str := ""
	fmt.Printf("%T\n", str)
	// str := ""
	// var str string
	// func AppendInt(dst []byte, i int64, base int) []byte
	str = strconv.AppendInt(str, 123, 10)
	str = strconv.AppendBool(str, false)

	fmt.Printf("%s\n", str)

	a := strconv.FormatBool(false)
	b := strconv.FormatInt(123, 10)
	fmt.Printf("%s %T, %s %T\n", a, a, b, b) //false string, 123 string

	c, _ := strconv.ParseInt("123", 10, 32)
	d, _ := strconv.ParseBool("true")
	fmt.Printf("%d %T, %t %T\n", c, c, d, d) //123 int64, true bool

	e, _ := strconv.Atoi("12354")
	fmt.Printf("%d\n", e)
}
