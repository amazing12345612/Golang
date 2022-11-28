/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
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

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list3 := new(ListNode)
	p := list3
	p1 := list1
	p2 := list2
	for {
		if p1 == nil {
			p.Next = p2
			break
		}
		if p2 == nil {
			p.Next = p1
			break
		}
		if p1.Val > p2.Val {
			p.Next = p2
			p = p.Next
			p2 = p2.Next
		} else {
			p.Next = p1
			p = p.Next
			p1 = p1.Next
		}
	}
	return list3.Next
}

// @lc code=end
