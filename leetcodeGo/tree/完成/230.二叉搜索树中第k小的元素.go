/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	left := Count(root.Left)

	if left == k-1 {
		return root.Val
	} else if left < k-1 {
		return kthSmallest(root.Right, k-1-left)
	} else {
		return kthSmallest(root.Left, k)
	}

}
func Count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Count(root.Left) + Count(root.Right) + 1
}

// @lc code=end
