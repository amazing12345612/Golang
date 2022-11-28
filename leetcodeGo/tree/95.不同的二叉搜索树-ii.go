/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) (result []*TreeNode) {
	l := make([]*TreeNode, n)
	for i := 1; i <= n; i++ {
		l = append(l, &TreeNode{Val: i})
	}
	var dfs func(list []*TreeNode) *TreeNode
	dfs = func(list []*TreeNode) *TreeNode {
		if len(list) == 0 {
			return nil
		}
		var root *TreeNode
		for i := 0; i < len(list); i++ {
			root.Val = list[i].Val
			root.Left = dfs(list[:i])
			root.Right = dfs(list[i+1:])
		}
		result = append(result, root)
		return root

	}
	dfs(l)
	return

}
func main() {
	a := []int{1, 2}
	fmt.Println(a[:0], len(a[:0]))
}

// @lc code=end
