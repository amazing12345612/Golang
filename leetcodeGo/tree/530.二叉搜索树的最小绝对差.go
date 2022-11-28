/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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

func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt16
	n := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		n = append(n, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	for i := 1; i < len(n); i++ {
		if n[i]-n[i-1] < min {
			min = n[i] - n[i-1]
		}
	}
	return min
}
func Abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -1
	}
}

// @lc code=end
