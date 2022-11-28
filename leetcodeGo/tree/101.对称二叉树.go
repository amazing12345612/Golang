/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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

import "fmt"

func isSymmetric(root *TreeNode) bool {

	r := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		r = append(r, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	fmt.Println(r)
	for i := 0; i != len(r)-1-i; i++ {
		if r[i] != r[len(r)-i-1] {
			return false
		}
	}
	return true
}
func isSymmetric1(root *TreeNode) bool {

	var dfs func(p *TreeNode, q *TreeNode) bool
	dfs = func(p *TreeNode, q *TreeNode) bool {
		if p != nil && q == nil || p == nil && q != nil {
			return false
		} else if p == nil && q == nil {
			return true
		} else {
			if p.Val != q.Val {
				return false
			} else {
				return dfs(p.Right, q.Left) && dfs(p.Left, q.Right)
			}
		}
	}
	if root != nil {
		return dfs(root.Left, root.Right)
	}
	return true
}

// @lc code=end
