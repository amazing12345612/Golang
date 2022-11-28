/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
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

func hasCycle(head *ListNode) bool {
	m := make(map[interface{}]int)
	for i := 0; head != nil; i++ {
		if m[head] >= 0 {
			return true
		}
		m[head] = i
		head = head.Next
	}
	return false
}

//快慢指针
func hasCycle1(head *ListNode) bool {
	fast := head
	flow := head
	for fast != nil && flow != nil {
		fast = fast.Next.Next
		flow = fast.Next
		if fast == flow {
			return true
		}
	}
	return false

}

// @lc code=end
