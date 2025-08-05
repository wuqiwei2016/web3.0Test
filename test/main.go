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

func init() {
	fmt.Println("main init")
}

func main() {
	// fmt.Println("main.main() method invoked")
	var s string = "Hello, world!"
	var bytes []byte = []byte(s)
	fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
	fmt.Println(string(bytes))
}
