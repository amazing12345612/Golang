/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 *
 */
package everyday

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
func middle(head *ListNode) *ListNode {
	slow, fast := head, head
	//要先判断Next是否为nil,不然Next.Next就无意义
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	p1, p2 := head1, head2
	for p1 != nil && p2 != nil {
		next1 := p1.Next
		p1.Next = p2
		p1 = next1
		next2 := p2.Next
		p2.Next = p1
		p2 = next2
	}
	return head1
}
func ReorderList(head *ListNode) {
	mid := middle(head)
	t := mid.Next
	mid.Next = nil
	mid = t
	l2 := ReverseList(mid)
	l1 := head
	head = merge(l1, l2)

}
func Show(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// @lc code=end
