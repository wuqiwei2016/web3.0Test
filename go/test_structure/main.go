package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := make(map[string]int)
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(2)

	go func() {
		for {
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()

	go func() {
		for {
			lock.Lock()
			m["a"]++
			fmt.Println(m["a"])
			lock.Unlock()
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout, stopping")
	}
}

// var b int = 10

// func main() {
// 	var b int = 4
// 	fmt.Println("local variable, b = ", b)
// 	if b := 3; b == 3 {
// 		fmt.Println("if statement, b = ", b)
// 		b--
// 		fmt.Println("if statement, b = ", b)
// 	}
// 	fmt.Println("local variable, b = ", b)
// 	fmt.Println("test = ", test())
// }

// func test() int {
// 	return b
// }

// type A struct {
// 	i int
// }

// func (a *A) add(v int) int {
// 	a.i += v
// 	return a.i
// }

// // 声明函数变量
// var function1 func(int) int

// // 声明闭包
// var squart2 func(int) int = func(p int) int {
// 	p *= p
// 	return p
// }

// func main() {
// 	a := A{1}
// 	// 把方法赋值给函数变量
// 	function1 = a.add

// 	// // 声明一个闭包并直接执行
// 	// // 此闭包返回值是另外一个闭包（带参闭包）
// 	// returnFunc := func() func(int, string) (int, string) {
// 	// 	fmt.Println("this is a anonymous function")
// 	// 	return func(i int, s string) (int, string) {
// 	// 		return i, s
// 	// 	}
// 	// }()

// 	// // 执行returnFunc闭包并传递参数
// 	// ret1, ret2 := returnFunc(1, "test")
// 	// fmt.Println("call closure function, return1 = ", ret1, "; return2 = ", ret2)

// 	fmt.Println("a.i = ", a.i)
// 	fmt.Println("after call function1, a.i = ", function1(1))
// 	fmt.Println("a.i = ", a.i)
// }

// var name1 = struct {
// 	file1  string `<tag1>:"<any string>"`
// 	file2  int    `<tag2>:"<any string>"`
// 	int    int
// 	string string
// }{
// 	file1:  "1",
// 	file2:  1,
// 	int:    2,
// 	string: "2",
// }

// var name2 = struct{}{}

// type A struct {
// 	a string
// }

// type B struct {
// 	A
// 	b string
// }

// type C struct {
// 	A
// 	B
// 	a string
// 	b string
// 	c string
// }

// type D struct {
// 	a A
// 	b B
// }

// func main() {
// 	a := A{a: "a"}
// 	b := B{A: a, b: "b"}
// 	c := C{A: a, B: b, a: "ca", b: "cb", c: "c"}
// 	fmt.Println(b.a)
// 	fmt.Println(c.A.a, c.a)
// 	fmt.Println(c.B.b, c.b)
// 	fmt.Println(c)
// 	// fmt.Println(name1.file1)
// 	// fmt.Println(name1.file2)
// 	// fmt.Println(name1.int)
// 	// fmt.Println(name1.string)

//		// var name4 = struct {
//		// 	file1 string
//		// 	file2 int
//		// 	int
//		// 	string
//		// }{
//		// 	file1: "2",
//		// 	file2: 555,
//		// 	int:   4,
//		// }
//		// fmt.Println(name4.file1)
//		// fmt.Println(name4.file2)
//		// fmt.Println(name4.int)
//		// fmt.Println(name4.string)
//		// fmt.Println(name2)
//	}
// type A struct {
// 	a string
// }

// func (a A) string() string {
// 	return a.a
// }

// func (a A) stringA() string {
// 	return a.a
// }

// func (a A) setA(v string) {
// 	a.a = v
// }

// func (a *A) stringPA() string {
// 	return a.a
// }

// func (a *A) setPA(v string) {
// 	a.a = v
// }

// type B struct {
// 	A
// 	b string
// }

// func (b B) string() string {
// 	return b.b
// }

// func (b B) stringB() string {
// 	return b.b
// }

// func value(a A, value string) {
// 	a.a = value
// }

// func point(a *A, value string) {
// 	a.a = value
// }

// type C struct {
// 	B
// 	a string
// 	b string
// 	c string
// 	d []byte
// }

// func (c C) string() string {
// 	return c.c
// }

// func (c C) modityD() {
// 	c.d[2] = 3
// }

// func callStructMethod() {
// 	var a A
// 	a = A{
// 		a: "a",
// 	}
// 	a.string()
// }

// func NewC() C {
// 	return C{
// 		B: B{
// 			A: A{
// 				a: "ba",
// 			},
// 			b: "b",
// 		},
// 		a: "ca",
// 		b: "cb",
// 		c: "c",
// 		d: []byte{1, 2, 3},
// 	}
// }

// func test1() {
// 	fmt.Println("test")
// }

// func main() {

// for {
// 	// 中断select
// 	select {
// 	case <-time.After(time.Second * 2):
// 		fmt.Println("过了2秒")
// 	case <-time.After(time.Second):
// 		fmt.Println("进过了1秒")
// 		if true {
// 			break
// 		}
// 		fmt.Println("break 之后")
// 	}
// }
// test1()
// a := A{
// 	a: "a",
// }

// ///////
// copy := a
// copy.a = "121212"
// fmt.Println(copy.a)
// fmt.Println(a.a)
// pa := &a
// a.setA("va")
// a.setPA("pa")
// pa.setA("pppp")
// pa.setPA("vvvv")
// fmt.Println(a.string())

// // a A
// fmt.Println(a.string())
// fmt.Println(a.stringA())
// fmt.Println(a.stringPA())
// fmt.Println(a.setA("1a"))

// fmt.Println(a.setPA("2a"))

// fmt.Println(pa.string())
// fmt.Println(pa.stringA())
// fmt.Println(pa.stringPA())
// a.stringPA()
// a.setPA("2a")
// a.stringPA()
// a.stringA()

// c := NewC()
// cp := &c
// fmt.Println(c.string())
// fmt.Println(c.stringA())
// fmt.Println(c.stringB())

// fmt.Println(cp.string())
// fmt.Println(cp.stringA())
// fmt.Println(cp.stringB())

// c.setA("1a")
// fmt.Println("------------------c.setA")
// fmt.Println(c.A.a)
// fmt.Println(cp.A.a)

// cp.setA("2a")
// fmt.Println("------------------cp.setA")
// fmt.Println(c.A.a)
// fmt.Println(cp.A.a)

// c.setPA("3a")
// fmt.Println("------------------c.setPA")
// fmt.Println(c.A.a)
// fmt.Println(cp.A.a)

// cp.setPA("4a")
// fmt.Println("------------------cp.setPA")
// fmt.Println(c.A.a)
// fmt.Println(cp.A.a)

// cp.modityD()
// fmt.Println("------------------cp.modityD")
// fmt.Println(cp.d)
// }
