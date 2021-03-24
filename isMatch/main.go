package main

import "fmt"

func isMatch(s string, p string) bool {
	if s == "" && p == "" {
		return true
	} else if s == "" || p == "" {
		return false
	}

	sp, pp := 0, 0
	var pre byte = 0
	for sp < len(s) && pp < len(p) {
		switch p[pp] {
		case '.':
			sp++
			pp++
		case '*':
			if pre == 0 {
				return true
			}
			for sp < len(s) && s[sp] == pre {
				sp++
			}
			pp++
		default:
			pre = p[pp]
			if pp + 1 < len(p) && p[pp+1] == '*' {
				pp++
				continue
			}

			if s[sp] != p[pp] {
				return false
			}
			sp++
			pp++
		}
	}

	if sp == len(s) && pp == len(p) {
		return true
	}
	return false
}

func main() {
	// fmt.Println(isMatch("", ""))	// true
	// fmt.Println(isMatch("", "1"))	// false
	// fmt.Println(isMatch("1", ""))	// false
	// fmt.Println(isMatch("aab", "c*a*b"))	// true
	// fmt.Println(isMatch("mississippi", "mis*is*p*."))	// false
	fmt.Println(isMatch("ab", ".*c"))
}