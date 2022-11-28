/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
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

func findDuplicateSubtrees(root *TreeNode) (a []*TreeNode) {
	t := make(map[string]*TreeNode)
	r := make(map[*TreeNode]interface{})
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return ""
		}
		s := fmt.Sprintf("%d(%s)(%s)", root.Val, root.Left, root.Right)
		if p, ok := t[s]; ok {
			r[p] = ""
		} else {
			t[s] = root
		}
		return s
	}
	dfs(root)
	for i := range r {
		a = append(a, i)
	}
	return
}

// @lc code=end
