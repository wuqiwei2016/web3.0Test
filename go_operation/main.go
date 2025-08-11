package main

import (
	"fmt"
)

// func main() {
// 	// a, b := 1, 2
// 	// sum := a + b
// 	// sub := a - b
// 	// mul := a * b
// 	// div := a / b
// 	// mod := a % b

// 	// fmt.Println(sum, sub, mul, div, mod)

// 	a := 10 + 0.1
// 	b := byte(1) + 1
// 	c := 1.01
// 	// fmt.Println(a, b)

// 	// sum := a + float64(b)
// 	// fmt.Println(sum)

// 	// sub := byte(a) - b
// 	// fmt.Println(sub)

// 	mul := a * c
// 	div := byte(a) / b

// 	fmt.Println(mul, div)

// }
func main() {
	// a, b := 1, 2
	// var c int
	// c = a + b
	// fmt.Println("c = a + b, c =", c)

	// plusAssignment(c, a)
	// subAssignment(c, a)
	// mulAssignment(c, a)
	// divAssignment(c, a)
	// modAssignment(c, a)
	// leftMoveAssignment(c, a)
	// rightMoveAssignment(c, a)
	// andAssignment(c, a)
	// orAssignment(c, a)
	// norAssignment(c, a)
	a := 4
	var ptr *int
	fmt.Println(a)

	ptr = &a
	fmt.Printf("*ptr 为 %d\n", *ptr)
	fmt.Printf("*ptr 为 %d\n", ptr)
}

func plusAssignment(c, a int) {
	c += a // c = c + a
	fmt.Println("c += a, c =", c)
}

func subAssignment(c, a int) {
	c -= a // c = c - a
	fmt.Println("c -= a, c =", c)
}

func mulAssignment(c, a int) {
	c *= a // c = c * a
	fmt.Println("c *= a, c =", c)
}

func divAssignment(c, a int) {
	c /= a // c = c / a
	fmt.Println("c /= a, c =", c)
}

func modAssignment(c, a int) {
	c %= a // c = c % a
	fmt.Println("c %= a, c =", c)
}

func leftMoveAssignment(c, a int) {
	c <<= a // c = c << a
	fmt.Println("c <<= a, c =", c)
}

func rightMoveAssignment(c, a int) {
	c >>= a // c = c >> a
	fmt.Println("c >>= a, c =", c)
}

func andAssignment(c, a int) {
	c &= a // c = c & a
	fmt.Println("c &= a, c =", c)
}

func orAssignment(c, a int) {
	c |= a // c = c | a
	fmt.Println("c |= a, c =", c)
}

func norAssignment(c, a int) {
	c ^= a // c = c ^ a
	fmt.Println("c ^= a, c =", c)
}
