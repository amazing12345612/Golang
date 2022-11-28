/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	q := []*TreeNode{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		q = append(q, root)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	for i := 0; i < len(q)-1; i++ {
		q[i].Right = q[i+1]
		q[i].Left = nil
	}
	root = q[0]
}

//递归
//思路就是把左子树变为链表，再把右子树变为链表，再拼接两个到root的右边

func flatten1(root *TreeNode) {
	if root == nil {
		return
	}
	flatten1(root.Left)
	root.Left = nil
	p := root.Right
	root.Right = root.Left
	flatten1(p)
	for root.Right != nil {
		root = root.Right
	}
	root.Right = p

}

// @lc code=end
