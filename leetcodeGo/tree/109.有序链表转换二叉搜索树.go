/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	r := []int{}
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}

	var dfs func(r []int) *TreeNode
	dfs = func(r []int) *TreeNode {
		if len(r) == 1 {
			t := &TreeNode{Val: r[0]}
			return t
		}
		mid := len(r) / 2
		if len(r) == 2 {
			root := new(TreeNode)
			root.Val = r[mid]
			root.Left = dfs(r[:mid])
			root.Right = nil
			return root
		}
		root := new(TreeNode)
		root.Val = r[mid]
		root.Left = dfs(r[:mid])
		root.Right = dfs(r[mid+1:])

		return root
	}
	root := dfs(r)
	return root

}

// @lc code=end
