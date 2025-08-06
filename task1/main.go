package main

import (
	"fmt"
	"strings"
)

// ... existing code ...
// func isValid(s string) bool {
// 	// ... existing code ...
// }

func isValidSimple(s string) bool {
	// 循环替换成对的括号，直到无法替换
	for {
		originalLen := len(s)
		// 替换所有成对的括号
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "{}", "")
		s = strings.ReplaceAll(s, "[]", "")
		// 输出长度
		fmt.Println(len(s))
		// 如果长度没变，说明没有可替换的括号了
		if len(s) == originalLen {
			break
		}
	}
	if s == "" {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println("1.有效括号()", isValidSimple("()"))
	// fmt.Println("2. 有效括号()[]{}", isValidSimple("()[]{}"))
	// fmt.Println("3. 有效括号(]", isValidSimple("(]"))
	// fmt.Println("4. 有效括号([])", isValidSimple("([])"))
	// fmt.Println("5. 有效括号([{)]})", isValidSimple("([{)]})"))
	// fmt.Println("6. 有效括号((", isValidSimple("(("))
	// fmt.Println("7. 有效括号(([]){})", isValidSimple("(([]){})"))
	// fmt.Println("8. 有效括号[({(())}[()])]", isValidSimple("[({(())}[()])]"))
}
