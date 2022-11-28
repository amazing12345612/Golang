/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(dfs(root.Left)-dfs(root.Right)) >= 1 && isBalanced(root.Left) && isBalanced(root.Right)

}
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(dfs(root.Left), dfs(root.Right)) + 1
}
func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
