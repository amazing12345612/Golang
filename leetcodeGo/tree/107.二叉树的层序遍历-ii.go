/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
	_ "fmt"
	_ "sort"
)

func levelOrderBottom(root *TreeNode) (result [][]int) {
	if root == nil {
		return [][]int{}
	}
	l := make([]struct{ x, y int }, 0)
	stake := make([]struct {
		x *TreeNode
		y int
	}, 0)
	stake = append(stake, struct {
		x *TreeNode
		y int
	}{root, 0})

	for len(stake) > 0 {
		r := stake[len(stake)-1]
		stake = stake[:len(stake)-1]
		l = append(l, struct {
			x int
			y int
		}{r.y, r.x.Val})
		if r.x.Right != nil {
			stake = append(stake, struct {
				x *TreeNode
				y int
			}{r.x.Right, r.y + 1})
		}
		if r.x.Left != nil {
			stake = append(stake, struct {
				x *TreeNode
				y int
			}{r.x.Left, r.y + 1})
		}
	}
	n := l[len(l)-1].x
	r := make([][]int, n+1)
	for _, v := range l {
		r[v.x] = append(r[v.x], v.y)
	}
	for i := len(r) - 1; i >= 0; i-- {
		result = append(result, r[i])
	}
	return result

}

// @lc code=end
