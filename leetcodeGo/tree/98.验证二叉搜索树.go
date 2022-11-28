/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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

import "math"

func isValidBST(root *TreeNode) bool {
	l := math.MinInt64
	r := math.MaxInt64
	var dfs func(root *TreeNode, l, r int) bool
	dfs = func(root *TreeNode, l, r int) bool {
		if root == nil {
			return true
		}
		if root.Left.Val <= l || root.Right.Val >= r {
			return false
		}
		return dfs(root.Left, l, root.Val) && dfs(root.Right, root.Val, r)
	}
	t := dfs(root, l, r)
	return t
}

// @lc code=end
