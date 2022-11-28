/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
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

import "fmt"

func rotateRight(head *ListNode, k int) *ListNode {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	reverse(nums[len(nums)-k:])
	fmt.Println(nums)
	reverse(nums[:len(nums)-k])
	fmt.Println(nums)
	reverse(nums)
	fmt.Println(nums)
	for i := 0; head != nil && i < len(nums); i++ {
		head.Val = nums[i]
	}
	return head
}
func reverse(nums []int) {
	m := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[i] = nums[len(nums)-i-1]
	}
	nums = m
}

// func rotateRight1(head *ListNode, k int) *ListNode {

// }

// @lc code=end
