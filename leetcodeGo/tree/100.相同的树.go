/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q == nil || p == nil && q != nil {
		return false
	} else {
		if p.Val != p.Val {
			return false
		} else {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}

}

// @lc code=end
