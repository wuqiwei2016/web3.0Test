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

// func init() {
// 	fmt.Println(s1)
// 	fmt.Println("main init")
// 	fmt.Println("main 11111")
// }

func main() {
	fmt.Println("main.main() method invoked")
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
