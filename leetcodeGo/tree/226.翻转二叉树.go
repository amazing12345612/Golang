/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	t := new(TreeNode)
	t = root.Left
	root.Left = root.Right
	root.Right = t
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// @lc code=end
