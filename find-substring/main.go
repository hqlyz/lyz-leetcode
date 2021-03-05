package main

import (
	"fmt"
	"strings"
)

// ---困难---
// 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

//

// 示例 1：

// 输入：
//   s = "barfoothefoobarman",
//   words = ["foo","bar"]
// 输出：[0,9]
// 解释：
// 从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
// 输出的顺序不重要, [9,0] 也是有效答案。
// 示例 2：

// 输入：
//   s = "wordgoodgoodgoodbestword",
//   words = ["word","good","best","word"]
// 输出：[]

func findSubstring(s string, words []string) []int {
	result := []int{}
	if len(s) == 0 || len(words) == 0 {
		return result
	}
	step := len(words[0])

	temp := []string{}
	test("", words, &temp)
	m := map[int]bool{}
	// find all "aa" in "aaa"
	for j := 0; j < step; j++ {
		for _, v := range temp {
			i := strings.Index(s, v)
			p := i
			for i >= 0 {
				if _, ok := m[p]; !ok {
					m[p] = true
					result = append(result, p)
				}
				p += step
				i = strings.Index(s[p:], v)
				// fmt.Println("current i:", i)
				// fmt.Println("current s:", s[p:])
				// time.Sleep(1 * time.Second)
			}
		}
	}
	return result
}

func test(prefix string, sl []string, result *[]string) {
	if len(sl) == 0 {
		*result = append(*result, prefix)
	}

	m := map[string]bool{}
	for k, v := range sl {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = true
		temp := []string{}
		temp = append(temp, sl[:k]...)
		temp = append(temp, sl[k+1:]...)
		test(prefix+v, temp, result)
	}
}

func main() {
	// words := []string{"word", "good", "best", "good"}
	words := []string{"aa", "aa"}
	s := "aaaaaaaaaaaaaa"
	fmt.Println(findSubstring(s, words))
}
