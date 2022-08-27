package main

import "fmt"

/*
继承  通过方法实现
封装  通过匿名字段实现
多态  通过接口实现
*/
// 继承
type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person //匿名字段  继承了Person 的字段
	id     int
	addr   string

	//name string //如果有同名字段，则使用该类的，继承的被覆盖
	// 若想使用 Person的name    s.Person.name 显示使用

	// int //非结构体匿名字段 使用方式  s.int

	//*Person //指针类型 初始化  Student{ &Person{"chen", 'm', 23}, 1, "sy"}
	// s2.Person = new(Person) 分配空间 再赋值
}

func test1() {
	var s1 Student = Student{Person{"chen", 'm', 23}, 1, "sy"}
	fmt.Println(s1)

	s2 := Student{Person{"chen", 'm', 23}, 2, "sy"}
	fmt.Printf("%+v", s2) // +v显示详细信息
	fmt.Println()
	// {{chen 109 23} 1 sy}
	// {Person:{name:chen sex:109 age:23} id:2 addr:sy}

	fmt.Println(s1.name)

	s1.addr = "bj"
	s1.age += 1
	fmt.Println(s1)

}
func main() {
	// test1()
	test2()
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func (p Point) DisplayPoint() {
	fmt.Printf("Point(X:%d,Y:%d)\n", p.X, p.Y)
}
func (c Circle) DisplayCircle() {
	fmt.Printf("Circle(X:%d,Y:%d,Radius:%d)\n", c.X, c.Y, c.Radius)
}

func (p *Point) SetPoint(x, y int) {
	p.X, p.Y = x, y
}
func test2() {
	p := Point{1, 2}
	// p.DisplayPoint()
	cir := Circle{p, 3}
	cir.DisplayPoint()
	cir.SetPoint(2, 3)
	cir.DisplayPoint()

	wh := Wheel{cir, 3}
	wh.DisplayCircle()
	wh.DisplayPoint()
}
