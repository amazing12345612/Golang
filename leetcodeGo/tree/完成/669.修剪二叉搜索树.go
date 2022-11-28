/*
 * @lc app=leetcode.cn id=669 lang=golang
 *
 * [669] 修剪二叉搜索树
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

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val >= low && root.Val <= high {
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
		return root
	}
	if root.Val < low {
		root.Right = trimBST(root.Right, low, high)
		return root.Right
	}
	if root.Val > high {
		root.Left = trimBST(root.Left, low, high)
		return root.Left
	}
	return root
}

// @lc code=end
