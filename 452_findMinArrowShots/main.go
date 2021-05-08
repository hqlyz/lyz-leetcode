package main

import (
	"log"
	"sort"
)

// medium
// input: [[10,16],[2,8],[1,6],[7,12]]
// output: 2

func main() {
	input := make([][]int, 0)
	input = append(input, []int{1, 2})
	input = append(input, []int{2, 3})
	input = append(input, []int{3, 4})
	input = append(input, []int{4, 5})
	log.Println(findMinArrowShots(input))
}

func findMinArrowShots(points [][]int) int {
	length := len(points)
	if length < 2 {
		return length
	}

	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	curXCoord := points[0][1]
	count := 1
	for i := 1; i < len(points); i++ {
		if curXCoord < points[i][0] {
			count++
			curXCoord = points[i][1]
		}
	}
	return count
}
