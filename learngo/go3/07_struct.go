package main

import "fmt"

type FuncType func(int, int) int

type Student struct {
	id     int
	Name   string
	gender byte
	age    int
	Addr   string
}

type Deng struct {
	s    Student
	flag bool
}

func (s *Student) String() string {
	return fmt.Sprintf("姓名:\t%s\n地址:\t%s", s.Name, s.Addr)
}

func main() {
	var s11 Student
	s11.age = 10
	s1 := Student{1, "zhang", 0, 25, "beijing"}
	fmt.Println(s1)

	s2 := Student{id: 2, Name: "wang"}
	fmt.Println(s2)

	var p *Student
	p = &s1
	fmt.Println(*p)

	//成员使用
	s := Student{}
	s.id = 3
	s.Name = "me"
	fmt.Println(s.Name)

	//指针

	p.id = 3
	// (*p).name = "cheng"
	p.Name = "cheng" //与上一行等价
	fmt.Println(*p)  //输出 ‘{内容}’
	fmt.Println(p)   //输出 ’&{内容}‘

	p2 := new(Student) //p2是一个指针
	p2.id = 4
	p2.Name = "he"
	fmt.Println(p2)

	//比较 和赋值 如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的
	s3 := s1
	fmt.Println(s3)
	fmt.Println(s3 == s1)
	fmt.Println(s2 == s1)

	s2 = s3
	fmt.Println(s2 == s3)

	d1 := Deng{s1, false}
	d2 := Deng{s1, true}
	fmt.Println("deng", d1 == d2)
	d2.flag = false
	fmt.Println("deng", d1 == d2)
	fmt.Println(d1, d2)

}
