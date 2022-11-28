/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode75

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	p := head
	count := 1
	for head.Next != nil {
		count++
		head = head.Next
	}
	i := count/2 + 1
	for i > 1 {
		p = p.Next
		i--
	}
	return p
}

// @lc code=end
