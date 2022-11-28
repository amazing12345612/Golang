/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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

import (
	"fmt"
)

func levelOrder(root *TreeNode) (ans [][]int) {
	var dfs func(root *TreeNode, t []int)
	dfs = func(root *TreeNode, t []int) {
		if root == nil {
			return
		}
		t = append(t, root.Val)
		a := []int{}
		dfs(root.Left, a)
		dfs(root.Right, a)
		ans = append(ans, a)
	}

	return

}

func levelOrder1(root *TreeNode) (ans [][]int) {
	r := []struct{ x, y int }{}
	var dfs func(root *TreeNode, n int)
	dfs = func(root *TreeNode, n int) {
		if root == nil {
			return
		}
		r = append(r, struct {
			x int
			y int
		}{root.Val, n})
		dfs(root.Left, n+1)
		dfs(root.Right, n+1)
	}
	dfs(root, 0)
	fmt.Println(r)
	t := make(map[int][]int)
	for _, v := range r {
		t[v.y] = append(t[v.y], v.x)
	}
	for i := 0; i < len(t); i++ {
		ans = append(ans, t[i])
	}
	return

}
func pop(q []*TreeNode) []*TreeNode {
	p := make([]*TreeNode, len(q)-1)
	for i := 0; i < len(q)-1; i++ {
		p[i] = q[i+1]
	}
	return p
}
func levelOrder2(root *TreeNode) (ans [][]int) {
	q := []*TreeNode{}
	if root != nil {
		q = append(q, root)
	}
	for len(q) > 0 {
		t := []int{}
		for i := 0; i < len(q); i++ {
			m := q[0]
			q = pop(q)
			fmt.Println(q)
			t = append(t, m.Val)
			if m.Left != nil {
				q = append(q, m.Left)
			}
			if m.Right != nil {
				q = append(q, m.Right)
			}
		}
		ans = append(ans, t)
	}
	return
}

// @lc code=end
