package main

import "fmt"

type FuncType func(int, int) int

type Student struct {
	id     int
	name   string
	gender byte
	age    int
	addr   string
}

func (s Student) String() string {
	return fmt.Sprintf("id\t:%d\nname\t:%s\n", s.id, s.name)
}

func main() {
	var s1 Student = Student{1, "chen", 0, 25, "sy"}
	fmt.Println(s1)

	s2 := Student{id: 2, name: "wang"}
	fmt.Println(s2)

	var p *Student
	p = &s1
	fmt.Println(*p)

	//成员使用
	s := Student{}
	s.id = 3
	s.name = "me"
	fmt.Println(s.name)

	//指针

	p.id = 3
	// (*p).name = "cheng"
	p.name = "cheng" //与上一行等价
	fmt.Println(*p)  //输出 ‘{内容}’
	fmt.Println(p)   //输出 ’&{内容}‘

	p2 := new(Student) //p2是一个指针
	p2.id = 4
	p2.name = "he"
	fmt.Println(p2)

	//比较 和赋值
	s3 := s1
	fmt.Println(s3)
	fmt.Println(s3 == s1)
	fmt.Println(s2 == s1)

	s2 = s3
	fmt.Println(s2 == s3)
}
