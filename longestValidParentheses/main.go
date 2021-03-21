package main

import "fmt"

// leetcode: 32
// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
// 示例 1：
// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"
// 示例 2：
// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"
// 示例 3：
// 输入：s = ""
// 输出：0
func longestValidParentheses(s string) int {
	count := 0
	maxCount := 0
	stack := []rune{}
	for _, ss := range s {
		switch ss {
		case '(':
			stack = append(stack, ss)
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack) - 1]
				count++
				if maxCount < count {
					maxCount = count
				}
			} else {
				count = 0
			}
		}
	}
	return maxCount * 2
}

func main() {
	s := "()(()"
	fmt.Println(longestValidParentheses(s))
	s = "(()"
	fmt.Println(longestValidParentheses(s))
	s = ")()())"
	fmt.Println(longestValidParentheses(s))
	s = ""
	fmt.Println(longestValidParentheses(s))
}