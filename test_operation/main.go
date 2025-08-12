package main

import "fmt"

func main() {
preset:
	for i := 0; i < 5; i++ {
		if i == 3 {
			break preset // 跳出标签标记的循环
		}
		fmt.Println("当前i的值为：", i)
	}
	// 	gotoPreset := false

	// preset:
	// 	a := 5

	// process:
	// 	if a > 0 {
	// 		a--
	// 		fmt.Println("当前a的值为：", a)
	// 		goto process
	// 	} else if a <= 0 {
	// 		// elseProcess:
	// 		if !gotoPreset {
	// 			gotoPreset = true
	// 			goto preset
	// 		} else {
	// 			goto post
	// 		}
	// 	}

	// post:
	//
	//	fmt.Println("main将结束，当前a的值为：", a)
}

// func main() {

// preset:
// 	a := 10

// 	// 使用标记
// outter:
// 	for i := 1; i <= 3; i++ {
// 		fmt.Printf("使用标记,外部循环, i = %d\n", i)
// 		for j := 5; j <= 10; j++ {
// 			fmt.Printf("使用标记,内部循环 j = %d\n", j)
// 			if j >= 7 {
// 				continue outter
// 			}
// 			fmt.Println("不使用标记，内部循环，在continue之后执行")
// 		}
// 	}
// // 方式1
// for i := 0; i < 10; i++ {
// 	fmt.Println("方式1，第", i+1, "次循环")
// }

// // 方式2
// b := 1
// for b < 10 {
// 	fmt.Println("方式2，第", b, "次循环")
// }

// // 方式3，无限循环
// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
// var started bool
// var stopped atomic.Bool
// for {
// 	if !started {
// 		started = true
// 		go func() {
// 			for {
// 				select {
// 				case <-ctx.Done():
// 					fmt.Println("ctx done")
// 					stopped.Store(true)
// 					return
// 				}
// 			}
// 		}()
// 	}
// 	fmt.Println("main")
// 	if stopped.Load() {
// 		break
// 	}
// }

// // 遍历数组
// var a [10]string
// a[0] = "Hello"
// for i := range a {
// 	fmt.Println("当前下标：", i)
// }
// for i, e := range a {
// 	fmt.Println("a[", i, "] = ", e)
// }

// // 遍历切片
// s := make([]string, 10)
// s[0] = "Hello"
// for i := range s {
// 	fmt.Println("当前下标：", i)
// }
// for i, e := range s {
// 	fmt.Println("s[", i, "] = ", e)
// }

// m := make(map[string]string)
// m["b"] = "Hello, b"
// m["a"] = "Hello, a"
// m["c"] = "Hello, c"
// for i := range m {
// 	fmt.Println("当前key：", i)
// }
// for k, v := range m {
// 	fmt.Println("m[", k, "] = ", v)
// }
// }

// type CustomType struct {
// }

// switch 语句
// func main() {
// a := "test string"

// // 1. 基本用法
// switch a {
// case "test":
// 	fmt.Println("a = ", a)
// case "s":
// 	fmt.Println("a = ", a)
// case "t", "test string": // 可以匹配多个值，只要一个满足条件即可
// 	fmt.Println("catch in a test, a = ", a)
// case "n":
// 	fmt.Println("a = not")
// default:
// 	fmt.Println("default case")
// }
// switch的使用
// a := "test string"
// switch a {
// case "test":
// 	fmt.Println(1111)
// case "s":
// 	fmt.Println(33333)
// case "t":
// 	fmt.Println("222222")
// default:
// 	fmt.Println(55555)
// }

// 变量b仅在当前switch代码块内有效
// switch b := 5; b {
// case 1:
// 	fmt.Println("b = 1")
// case 2:
// 	fmt.Println("b = 2")
// case 3, 4:
// 	fmt.Println("b = 3 or 4")
// case 5:
// 	fmt.Println("b = 5")
// default:
// 	fmt.Println("b = ", b)
// }

// switch b := 5; b {
// case 5:
// 	fmt.Println(111)
// case 3, 4:
// 	fmt.Println(2222)
// default:
// 	fmt.Println(33333)
// }
// a := "test string"
// // 不指定判断变量，直接在case中添加判定条件
// b := 5
// switch {
// case a == "t":
// 	fmt.Println("a = t")
// case b == 3:
// 	fmt.Println("b = 5")
// case b == 5, a == "test string":
// 	fmt.Println("a = test string; or b = 5")
// default:
// 	fmt.Println("default case")
// }
// 不指定判断变量，直接在case中添加判定条件
// switch {
// case a == "test":
// 	fmt.Println(1111)
// case b == 5:
// 	fmt.Println(2222)
// }

// var d interface{}
// // var e byte = 1
// e := CustomType{}
// d = &e
// switch t := d.(type) {
// case byte:
// 	fmt.Println("d is byte type, ", t)
// case *byte:
// 	fmt.Println("d is byte point type, ", t)
// case *int:
// 	fmt.Println("d is int type, ", t)
// case *string:
// 	fmt.Println("d is string type, ", t)
// case *CustomType:
// 	fmt.Println("d is CustomType pointer type, ", t)
// case CustomType:
// 	fmt.Println("d is CustomType type, ", t)
// default:
// 	fmt.Println("d is unknown type, ", t)
// }
// }

// func main() {
// 	// 声明赋值语句
// 	if ok := method(); ok {
// 		// success
// 		fmt.Println("success")
// 	} else if err := methodThrowError(); err != nil {
// 		// error
// 	}

// var a int = 10
// if b := 2; a > 10 {
// 	b = 2
// 	// c = 2
// 	fmt.Println("a > 10")
// } else if c := 3; b > 1 {
// 	b = 3
// 	fmt.Println("b > 1")
// } else {
// 	fmt.Println("其他")
// 	if c == 3 {
// 		fmt.Println("c == 3")
// 	}
// 	fmt.Println(b)
// 	fmt.Println(c)
// }
// fmt.Println(a)
// }

// func method() bool {
// 	return true
// }

// func methodThrowError() error {
// 	return errors.New("error")
// }

// 运算符内容
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

// 运算符的内容
// func main() {
// 	// a, b := 1, 2
// 	// var c int
// 	// c = a + b
// 	// fmt.Println("c = a + b, c =", c)

// 	// plusAssignment(c, a)
// 	// subAssignment(c, a)
// 	// mulAssignment(c, a)
// 	// divAssignment(c, a)
// 	// modAssignment(c, a)
// 	// leftMoveAssignment(c, a)
// 	// rightMoveAssignment(c, a)
// 	// andAssignment(c, a)
// 	// orAssignment(c, a)
// 	// norAssignment(c, a)
// 	a := 4
// 	var ptr *int
// 	fmt.Println(a)
// 	ptr = &a
// 	fmt.Printf("*ptr 为 %d\n", *ptr)
// 	fmt.Printf("*ptr 为 %d\n", ptr)
// }

// func plusAssignment(c, a int) {
// 	c += a // c = c + a
// 	fmt.Println("c += a, c =", c)
// }

// func subAssignment(c, a int) {
// 	c -= a // c = c - a
// 	fmt.Println("c -= a, c =", c)
// }

// func mulAssignment(c, a int) {
// 	c *= a // c = c * a
// 	fmt.Println("c *= a, c =", c)
// }

// func divAssignment(c, a int) {
// 	c /= a // c = c / a
// 	fmt.Println("c /= a, c =", c)
// }

// func modAssignment(c, a int) {
// 	c %= a // c = c % a
// 	fmt.Println("c %= a, c =", c)
// }

// func leftMoveAssignment(c, a int) {
// 	c <<= a // c = c << a
// 	fmt.Println("c <<= a, c =", c)
// }

// func rightMoveAssignment(c, a int) {
// 	c >>= a // c = c >> a
// 	fmt.Println("c >>= a, c =", c)
// }

// func andAssignment(c, a int) {
// 	c &= a // c = c & a
// 	fmt.Println("c &= a, c =", c)
// }

// func orAssignment(c, a int) {
// 	c |= a // c = c | a
// 	fmt.Println("c |= a, c =", c)
// }

// func norAssignment(c, a int) {
// 	c ^= a // c = c ^ a
// 	fmt.Println("c ^= a, c =", c)
// }
