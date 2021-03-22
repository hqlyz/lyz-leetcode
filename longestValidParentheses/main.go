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
	maxLen := 0
	dp := make([]int, len(s))
	for k, v := range s {
		if v == ')' && k > 0 {
			if s[k-1] == '(' {
				if k < 2 {
					dp[k] = 2
				} else {
					dp[k] = 2 + dp[k-2]
				}
			} else {
				if (k-dp[k-1]-1) >= 0 && s[k-dp[k-1]-1] == '(' {
					if k-dp[k-1]-2 >= 0 {
						dp[k] = dp[k-1] + 2 + dp[k-dp[k-1]-2]
					} else {
						dp[k] = dp[k-1] + 2
					}
				}
			}
			maxLen = max(maxLen, dp[k])
		}
	}

	return maxLen
}

func max(maxLen, i int) int {
	if maxLen < i {
		return i
	}
	return maxLen
}

func main() {
	s := "()(()"
	fmt.Println(longestValidParentheses(s))
	s = "(()"
	fmt.Println(longestValidParentheses(s))
	s = "())"
	fmt.Println(longestValidParentheses(s))
	s = ")()())"
	fmt.Println(longestValidParentheses(s))
	s = ""
	fmt.Println(longestValidParentheses(s))
}
