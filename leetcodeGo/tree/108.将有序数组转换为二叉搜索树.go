/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
package tree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		n := &TreeNode{Val: nums[0]}
		return n
	}
	n := len(nums) / 2
	root := &TreeNode{Val: nums[n]}
	if len(nums) == 2 {
		root.Left = sortedArrayToBST(nums[:n])
		root.Right = nil
		return root
	}
	root.Left = sortedArrayToBST(nums[:n])
	root.Right = sortedArrayToBST(nums[n+1:])
	return root

}

// @lc code=end
