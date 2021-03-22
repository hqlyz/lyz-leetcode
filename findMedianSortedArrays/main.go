package main

import "fmt"

// leetcode: 4
// hard
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 示例 1：
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：
// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
// 示例 3：
// 输入：nums1 = [0,0], nums2 = [0,0]
// 输出：0.00000
// 示例 4：
// 输入：nums1 = [], nums2 = [1]
// 输出：1.00000
// 示例 5：
// 输入：nums1 = [2], nums2 = []
// 输出：2.00000
// O(log (m+n))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := []int{}

	p1 := len(nums1)
	p2 := len(nums2)
	index := (p1 + p2) / 2
	l1, h1 := 0, p1
	l2, h2 := 0, p2

	for len(result) <= index {
		if p1 > 0 {
			for l1 < h1 && l2 < h2 && nums1[h1 - 1] > nums2[l2] {
				h1 = (l1 + h1) / 2
			}
			if l1 < h1 || (l2 < h2 && nums1[h1] <= nums2[l2]) {
				result = append(result, nums1[l1:h1]...)
				if l1 < h1 {
					l1 = h1
				} else {
					l1++
				}
			} 
			h1 = p1
		}

		if p2 > 0 {
			for l1 < h1 && l2 < h2 && nums2[h2 - 1] > nums1[l1] {
				h2 = (l2 + h2) / 2
			}
			if l2 < h2 || (l1 < h1 && nums2[h2] <= nums1[l1]) {
				result = append(result, nums2[l2:h2]...)
				if l2 < h2 {
					l2 = h2
				} else {
					l2++
				}
			}
			h2 = p2
		}
	}
	
	if (p1 + p2) % 2 == 0 {
		return (float64(result[index - 1]) + float64(result[index])) / 2
	}
	return float64(result[index])
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-1, 3}
	fmt.Printf("%.5f\n", findMedianSortedArrays(nums1, nums2))
	// nums1 = []int{1, 2}
	// nums2 = []int{3, 4}
	// fmt.Printf("%.5f\n", findMedianSortedArrays(nums1, nums2))
	// nums1 = []int{0, 0}
	// nums2 = []int{0, 0}
	// fmt.Printf("%.5f\n", findMedianSortedArrays(nums1, nums2))
	// nums1 = []int{}
	// nums2 = []int{1}
	// fmt.Printf("%.5f\n", findMedianSortedArrays(nums1, nums2))
	// nums1 = []int{2}
	// nums2 = []int{}
	// fmt.Printf("%.5f\n", findMedianSortedArrays(nums1, nums2))
}