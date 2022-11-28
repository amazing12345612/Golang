/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
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
package everyday

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	stake := []*TreeNode{}
	m := []struct{ x, y int }{{0, 0}}
	l := math.MinInt32
	p := root
	for len(stake) > 0 || p != nil {
		if p != nil {
			stake = append(stake, p)
			p = p.Left
		} else {
			p = stake[len(stake)-1]
			if m[0].x == p.Val {
				m[0].y++
			} else {
				m[0].x = p.Val
				m[0].y = 0
			}
			if m[0].y > l {
				l = m[0].y
			}
			stake = stake[:len(stake)-1]
			p = p.Right

		}
	}
	return l
}

// @lc code=end
