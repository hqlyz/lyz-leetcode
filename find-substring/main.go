package main

import (
	"fmt"
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

	m := map[string]int{}

	for _, word := range words {
		incWordCount(word, &m)
	}

	wordNum := len(words)
	wordLen := len(words[0])
	count := 0
	for i := 0; i < wordLen; i++ {
		count = 0
		left := i
		right := left
		tempM := map[string]int{}
		word := ""
		for right < len(s) {
			tempLeft := right
			right += wordLen
			if right > len(s) {
				right = len(s)
			}
			count++
			word = s[tempLeft:right]
			incWordCount(word, &tempM)
			for getWordCount(word, tempM) > getWordCount(word, m) {
				if getWordCount(word, m) == 0 {
					tempM = map[string]int{}
					count = 0
					left = right
					continue
				}
				tempLeft := left
				left += wordLen
				if left > len(s) {
					left = len(s)
				}
				decWordCount(s[tempLeft:left], &tempM)
				count--
			}
			if count == wordNum {
				result = append(result, left)
			}
		}
	}
	return result
}

func getWordCount(word string, words map[string]int) int {
	if _, ok := words[word]; !ok {
		return 0
	}
	return words[word]
}

func incWordCount(word string, words *map[string]int) {
	if _, ok := (*words)[word]; !ok {
		(*words)[word] = 1
		return
	}
	(*words)[word]++
}

func decWordCount(word string, words *map[string]int) {
	(*words)[word]--
}

func main() {
	// words := []string{"word", "good", "best", "good"}
	words := []string{"foo", "bar"}
	s := "barfoothefoobarman"
	fmt.Println(findSubstring(s, words))
}
