package main

import "fmt"

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// s 第 i 个字符和 p 第 j 个字符是否匹配
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if match(i, j-1, s, p) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				if match(i, j, s, p) {
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}
			}
		}
	}

	return dp[m][n]
}

func match(i, j int, s, p string) bool {
	if i == 0 {
		return false
	}

	if p[j-1] == '.' {
		return true
	}

	return s[i-1] == p[j-1]
}

func main() {
	fmt.Println(isMatch("", ""))                      // true
	fmt.Println(isMatch("", "1"))                     // false
	fmt.Println(isMatch("1", ""))                     // false
	fmt.Println(isMatch("aab", "c*a*b"))              // true
	fmt.Println(isMatch("mississippi", "mis*is*p*.")) // false
	fmt.Println(isMatch("abc", ".*c"))                // true
	fmt.Println(isMatch("ab", ".*c"))                 // false
}
