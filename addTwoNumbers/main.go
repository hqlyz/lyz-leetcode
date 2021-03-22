package main

import "fmt"

// leetcode: 2
// medium
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 示例 1：
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 示例 2：
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
// 示例 3：
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	result := &ListNode{}
	p := result
	addon := 0
	sum := 0
	for l1 != nil || l2 != nil {
		p.Next = &ListNode{}
		p = p.Next
		sum = getListNodeValue(l1) + getListNodeValue(l2) + addon
		addon = 0
		if sum >= 10 {
			addon = 1
			sum = sum % 10
		}
		p.Val = sum
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}	

	if addon == 1 {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = 1
	}

	return result.Next
}

func getListNodeValue(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}

// ListNode object
type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: &ListNode{Val: 6}}}
	l2 := &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8}}}
	printListNode(addTwoNumbers(l1, l2))
}

func printListNode(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
}