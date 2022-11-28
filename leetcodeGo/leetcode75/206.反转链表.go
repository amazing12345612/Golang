/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
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

func ReverseList(head *ListNode) *ListNode {
	r := []int{}
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	Newhead := new(ListNode)
	p := Newhead
	for i := len(r) - 1; i >= 0; i-- {
		temp := new(ListNode)
		temp.Val = r[i]
		p.Next = temp
		p = p.Next
	}
	return Newhead.Next
}
func ReverseList1(head *ListNode) *ListNode {
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
func ReverseList2(head *ListNode) *ListNode {
	//找到最后一个节点
	//1->2->3->4->5
	if head == nil || head.Next == nil {
		return head
	}
	//递归直到找到最后一个节点
	//5
	newhead := ReverseList2(head.Next)
	//1->2->3->4->5
	//5->next =4
	//1->2->3->4<->5
	head.Next.Next = head

	//断开4的指针
	head.Next = nil

	return newhead

}

// @lc code=end
