/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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

func countNodes(root *TreeNode) (r int) {
	queue := []*TreeNode{}
	queue = append(queue, root)
	count := 0
	for len(queue) > 0 {
		size := len(queue)
		if size != 1<<count {
			return (2 << count) - 1 + size
		}
		for i := 0; i < size; i++ {
			root := queue[0]
			queue = queue[1:]

			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		count++
	}
	return
}
func countNodes1(root *TreeNode) (r int) {
	if root == nil {
		return 0
	}
	left := countNodes1(root.Left)
	right := countNodes1(root.Right)
	if left == right {
		return countNodes1(root.Right) + 1<<left
	} else {
		return countNodes1(root.Left) + 1<<right
	}
}
func Height(root *TreeNode) int {
	count := 0
	for root != nil {
		root = root.Left
		count++
	}
	return count
}
func main() {
	fmt.Println(1<<0, 1<<1)
}

// @lc code=end
