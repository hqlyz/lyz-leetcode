package main

import (
	"log"
	"sort"
)

// input: [1, 2], [1, 2, 3]
// output: 2

func main() {
	children := []int{1, 2}
	cookies := []int{2, 1, 3}

	count := assignCookies(children, cookies)
	log.Println("count: ", count)
}

func assignCookies(children []int, cookies []int) int {
	sort.Slice(children, func(i, j int) bool {return children[i] < children[j]});
	sort.Slice(cookies, func(i, j int) bool {return cookies[i] < cookies[j]});

	count := 0
	for(len(children) > 0 && len(cookies) > 0) {
		if(cookies[0] >= children[0]) {
			children = children[1:]
			count++
		}
		cookies = cookies[1:]
	}
	return count
}