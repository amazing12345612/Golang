/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p1 := head
	p2 := head
	if p1.Next != nil {
		p2 = p1.Next
	} else {
		return p1
	}
	t1 := head
	t2 := head.Next
	for p1.Next != nil && p2.Next != nil {
		p1.Next = p2.Next
		p1 = p1.Next
		if p1.Next == nil {
			break
		}
		p2.Next = p1.Next
		p2 = p2.Next

	}
	p1.Next = t2
	p2.Next = nil
	return t1

}
func show(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}
func main() {
	head := new(ListNode)
	top := head
	head.Val = 0
	for i := 1; i < 5; i++ {
		head.Next = &ListNode{Val: i}
		head = head.Next
	}
	show(top)
	oddEvenList(top)
	show(top)

}

// @lc code=end
