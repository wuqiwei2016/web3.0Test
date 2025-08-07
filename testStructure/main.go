package main

import "fmt"

var name1 = struct {
	file1  string `<tag1>:"<any string>"`
	file2  int    `<tag2>:"<any string>"`
	int    int
	string string
}{
	file1:  "1",
	file2:  1,
	int:    2,
	string: "2",
}

var name2 = struct{}{}

type A struct {
	a string
}

type B struct {
	A
	b string
}

type C struct {
	A
	B
	a string
	b string
	c string
}

type D struct {
	a A
	b B
}

func main() {
	a := A{a: "a"}
	b := B{A: a, b: "b"}
	c := C{A: a, B: b, a: "ca", b: "cb", c: "c"}
	fmt.Println(b.a)
	fmt.Println(c.A.a, c.a)
	fmt.Println(c.B.b, c.b)
	fmt.Println(c)
	// fmt.Println(name1.file1)
	// fmt.Println(name1.file2)
	// fmt.Println(name1.int)
	// fmt.Println(name1.string)

	// var name4 = struct {
	// 	file1 string
	// 	file2 int
	// 	int
	// 	string
	// }{
	// 	file1: "2",
	// 	file2: 555,
	// 	int:   4,
	// }
	// fmt.Println(name4.file1)
	// fmt.Println(name4.file2)
	// fmt.Println(name4.int)
	// fmt.Println(name4.string)
	// fmt.Println(name2)
}
