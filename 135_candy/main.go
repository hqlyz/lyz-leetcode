package main

import "log"

// hard
// input: [1, 0, 2]
// output: 5

func main() {
	children := []int{1, 0, 2}
	log.Println("The smallest candies: ", candy(children))
}

func candy(children []int) int {
	if len(children) == 0 {
		return 0
	}
	candies := make([]int, len(children))
	for k := range candies {
		candies[k] = 1
	}

	for i := 0; i < len(children) - 1; i++ {
		if children[i+1] > children[i] {
			candies[i+1] = candies[i] + 1
		}
	}

	for i := len(children) - 1; i > 0; i-- {
		if children[i - 1] > children[i] && candies[i - 1] <= candies[i] {
			candies[i - 1] = candies[i] + 1
		}
	}

	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum
}