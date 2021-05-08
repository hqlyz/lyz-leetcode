package main

import (
	"log"
	"sort"
)

// medium
// input:  [[1,2], [2,4], [1,3]]
// output: 1
func main() {
	input := make([][]int, 0)
	input = append(input, []int{1, 2})
	input = append(input, []int{2, 4})
	input = append(input, []int{1, 3})
	log.Println(eraseOverlapIntervals(input))
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	count := 0
	index := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[index][1] > intervals[i][0] {
			// overlap
			count++
			continue
		}
		index = i
	}
	return count
}
