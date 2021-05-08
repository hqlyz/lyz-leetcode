package main

import "log"

// easy
// input: [1,0,0,0,1,0,0], 2
// output: true

func main() {
	input := []int{1,0,0,0,1,0,0}
	log.Println(canPlaceFlowers(input, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	temp := []int{0}
	temp = append(temp, flowerbed...)
	temp = append(temp, 0)
	for i := 1; i < len(temp) - 1; i++ {
		if temp[i] == 1 {
			continue
		}
		if temp[i - 1] == 0 && temp[i + 1] == 0 {
			temp[i] = 1
			n--
		}
		if n <= 0 {
			break
		}
	}
	return n <= 0
}
