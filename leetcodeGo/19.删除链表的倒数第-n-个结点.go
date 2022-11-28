/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	stake := []*ListNode{}
	tmp := &ListNode{Val: 0, Next: head}
	p := tmp
	for p != nil {
		stake = append(stake, p)
		p = p.Next
	}
	for len(stake) > 0 && n > 0 {
		n--
		stake = stake[:len(stake)-1]
	}
	t := stake[len(stake)-1]
	t.Next = t.Next.Next
	return tmp.Next

}

// @lc code=end
