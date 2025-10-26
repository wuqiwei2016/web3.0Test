package main

import (
	"fmt"
	"strings"
)

// 字符串处理、栈的使用
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

// 处理公共前缀
func dealCommonPreFix(strs []string) string {
	// 如果长度为空,则返回”“
	if len(strs) == 0 {
		return ""
	}
	// 初始化前缀
	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			// 前缀减少一个字符
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// plusOne 接收表示大整数的数字数组，返回加1后的结果数组
func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始处理
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		// 如果当前位加1后小于10，无需进位，直接返回
		if digits[i] < 10 {
			return digits
		}
		// 否则设置为0并继续向前处理进位
		digits[i] = 0
	}

	// 如果所有位都进位了（如输入[9,9]），需要在开头添加1
	return append([]int{1}, digits...)
}
func main() {
	// 字符串处理、栈的使用
	fmt.Println("1.有效括号()", isValidSimple("()"))
	fmt.Println("2. 有效括号()[]{}", isValidSimple("()[]{}"))
	fmt.Println("3. 有效括号(]", isValidSimple("(]"))
	fmt.Println("4. 有效括号([])", isValidSimple("([])"))
	fmt.Println("5. 有效括号([{)]})", isValidSimple("([{)]})"))
	fmt.Println("6. 有效括号((", isValidSimple("(("))
	fmt.Println("7. 有效括号(([]){})", isValidSimple("(([]){})"))
	fmt.Println("8. 有效括号[({(())}[()])]", isValidSimple("[({(())}[()])]"))
	// 公共前缀
	// 示例 1: 存在公共前缀
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Printf("输入: strs = %v\n输出: %s\n", strs1, dealCommonPreFix(strs1))
	strs2 := []string{"name", "nam", "a"}
	fmt.Printf("输入: strs = %v\n输出: %s\n", strs2, dealCommonPreFix(strs2))
	strs3 := []string{"apple", "app", "bann"}
	fmt.Printf("输入: strs = %v\n输出: %s\n", strs3, dealCommonPreFix(strs3))
	// 基本值类型
	fmt.Println(plusOne([]int{1, 2, 3}))    // 输出: [1 2 4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // 输出: [4 3 2 2]
	fmt.Println(plusOne([]int{9}))          // 输出: [1 0]
	fmt.Println(plusOne([]int{9, 9, 9}))    // 输出: [1 0 0 0]
}
