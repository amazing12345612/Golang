/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	Max := math.MinInt16

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := maxPathSum(root.Left)
		r := maxPathSum(root.Right)
		t1 := max(l, 0) + max(r, 0) + root.Val
		t2 := max(0, max(l, r)) + root.Val

		Max = max(Max, max(t1, t2))

		return t2
	}
	dfs(root)
	return Max

}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
