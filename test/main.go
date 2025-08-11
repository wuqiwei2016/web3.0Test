package main

// import "fmt"

// func main() {
// 	var s string = "Go语言"
// 	var bytes []byte = []byte(s)
// 	var runes []rune = []rune(s)

// 	fmt.Println("string length: ", len(s))
// 	fmt.Println("bytes length: ", len(bytes))
// 	fmt.Println("runes length: ", len(runes))
// }

//... existing code ...
import (
	"fmt"
)

// const name string = "Go语言"

// const c, b = 1, "2"

// const e, f = true, false

// type Gender string

// const (
//    Male   Gender = "Male"
//    Female Gender = "Female"
// )
// var s1 string = "Hello"
// var zero int
// var b1 = true

// var (
// 	i  int = 123
// 	b2 bool
// 	s2 = "test"
// )

// var (
// 	group = 2
// )

//	func init() {
//		fmt.Println(s1)
//		fmt.Println("main init")
//		fmt.Println("main 11111")
//	}
//
// src/net/http/server.go
// type ConnState int

// const (
// 	StateNew ConnState = iota
// 	StateActive
// 	StateIdle
// 	StateHijacked
// 	StateClosed
// )

// // src/time/time.go
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
const (
	a int = iota
	b
	c
	d
)

// type Gender byte

// const (
// 	Male Gender = iota
// 	Female
// )

// func (g *Gender) isMale() bool {
// 	return *g == Male
// }

// func (g *Gender) isFemale() bool {
// 	return *g == Female
// }

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// var Gender = Male
	// if Gender.isMale() {
	// 	fmt.Println("male")
	// } else {
	// 	fmt.Println("female")
	// }

	// fmt.Println(January)
	// fmt.Println(February)

	// fmt.Println("main.main() method invoked")
	// var s string = "Hello, world!"
	// var bytes []byte = []byte(s)
	// fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
	// fmt.Println(string(bytes))
}

//	func getMainVar() string {
//		fmt.Println("main.mainVar has been initialized")
//		return mainName
//	}
func method1() {
	// 方式1，类型推导，用得最多
	a := 1
	// 方式2，完整的变量声明写法
	var b int = 2
	// 方式3，仅声明变量，但是不赋值，
	var c int
	fmt.Println(a, b, c)
}

// 方式4，直接在返回值中声明
func method2() (a int, b string) {
	// 这种方式必须声明return关键字
	// 并且同样不需要使用，并且也不用必须给这种变量赋值
	return 1, "test"
}

func method3() (a int, b string) {
	a = 1
	b = "test"
	return
}

func method4() (a int, b string) {
	return
}
