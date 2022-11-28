/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 *
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历 迭代
//Stake 先向Left左节点不断遍历，遇到空节点，弹出处理，再向右节点
//寻找

func InorderTraversal(root *TreeNode) []int {
	stake := []*TreeNode{}
	r := []int{}
	for root != nil || len(stake) != 0 {
		if root != nil {
			stake = append(stake, root)
			root = root.Left
		} else {
			r = append(r, stake[len(stake)-1].Val)
			stake = stake[:len(stake)-1]
			root = root.Right
		}
	}
	return r
}

// @lc code=end
