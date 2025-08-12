package main

import "fmt"

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

// 	// var name4 = struct {
// 	// 	file1 string
// 	// 	file2 int
// 	// 	int
// 	// 	string
// 	// }{
// 	// 	file1: "2",
// 	// 	file2: 555,
// 	// 	int:   4,
// 	// }
// 	// fmt.Println(name4.file1)
// 	// fmt.Println(name4.file2)
// 	// fmt.Println(name4.int)
// 	// fmt.Println(name4.string)
// 	// fmt.Println(name2)
// }
type A struct {
	a string
}

func (a A) string() string {
	return a.a
}

func (a A) stringA() string {
	return a.a
}

func (a A) setA(v string) {
	a.a = v
}

func (a *A) stringPA() string {
	return a.a
}

func (a *A) setPA(v string) {
	a.a = v
}

type B struct {
	A
	b string
}

func (b B) string() string {
	return b.b
}

func (b B) stringB() string {
	return b.b
}

func value(a A, value string) {
	a.a = value
}

func point(a *A, value string) {
	a.a = value
}

type C struct {
	B
	a string
	b string
	c string
	d []byte
}

func (c C) string() string {
	return c.c
}

func (c C) modityD() {
	c.d[2] = 3
}

func callStructMethod() {
	var a A
	a = A{
		a: "a",
	}
	a.string()
}

func NewC() C {
	return C{
		B: B{
			A: A{
				a: "ba",
			},
			b: "b",
		},
		a: "ca",
		b: "cb",
		c: "c",
		d: []byte{1, 2, 3},
	}
}

func main() {

	a := A{
		a: "a",
	}
	copy := a
	copy.a = "copy"
	value(a, "any")
	point(&a, "any")

	pa := &a
	a.setA("va")
	a.setPA("pa")
	// pa.setA("pppp")
	// pa.setPA("vvvv")
	fmt.Println(a.string())

	// a A
	fmt.Println(a.string())
	fmt.Println(a.stringA())
	fmt.Println(a.stringPA())
	// fmt.Println(a.setA("1a"))

	// fmt.Println(a.setPA("2a"))

	fmt.Println(pa.string())
	fmt.Println(pa.stringA())
	fmt.Println(pa.stringPA())
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
}
