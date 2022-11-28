/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	queue := []*TreeNode{}
	sum := 0
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			root := queue[0]
			queue = queue[1:]
			if root.Left != nil {
				root.Left.Val += 10 * root.Val
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				root.Right.Val += 10 * root.Val
				queue = append(queue, root.Right)
			}
			if root.Left == nil && root.Right == nil {
				sum += root.Val
			}
		}
	}
	return sum
}

// @lc code=end
