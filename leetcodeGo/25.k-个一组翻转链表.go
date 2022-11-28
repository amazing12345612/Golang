/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
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

/*
	采用栈
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := &ListNode{Next: head}
	if head == nil {
		return nil
	}
	p1 := head
	//建立栈
	stake := []*ListNode{}

	for p1 != nil {
		//记录每次入栈的节点数量
		count := 0
		//head为该次入栈时头节点,t用于记录节点处理位置
		t := head
		for count < k && t != nil {
			stake = append(stake, t)
			count++
			//t最终指向的节点是下一次 入栈操作的头节点
			t = t.Next
		}
		//如果节点数量小于k,那么pre等于当前轮开始节点，后续节点就不处理
		if count < k {
			pre.Next = head
			break
		}
		//弹出栈中元素
		l := len(stake)
		for i := 0; i < l; i++ {
			n := stake[len(stake)-1]
			stake = stake[:len(stake)-1]
			//连接节点，pre指向下个节点
			pre.Next, pre = n, n
		}
		//连接下一个头节点
		pre.Next = t
		//head也指向下次头节点
		head = t

	}
	return pre.Next
}

// @lc code=end
