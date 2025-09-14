package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
	Perimeter() float64
}

/*
*
长方形
*/
type Rectangle struct {
	Width  float64
	Height float64
}

// 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

/*
*
圆
*/
type Circle struct {
	Radius float64
}

// 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:\n姓名: %s\n年龄: %d\n员工ID: %s\n",
		e.Name, e.Age, e.EmployeeID)
}

func main() {
	// 创建 Rectangle 实例
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("矩形 - 面积: %.2f, 周长: %.2f\n", rect.Area(), rect.Perimeter())

	// 创建 Circle 实例
	circle := Circle{Radius: 4}
	fmt.Printf("圆形 - 面积: %.2f, 周长: %.2f\n", circle.Area(), circle.Perimeter())

	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "E1001",
	}

	// 调用 PrintInfo 方法
	emp.PrintInfo()
}
