/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	stake := []*TreeNode{}
	var p1, p2, pre *TreeNode
	for len(stake) > 0 {
		for root != nil {
			stake = append(stake, root.Left)
			root = root.Left
		}
		root = stake[len(stake)-1]
		stake = stake[:len(stake)-1]
		if pre != nil && pre.Val > root.Val {
			p1 = root
			if p2 == nil {
				p2 = pre
			} else {
				break
			}
		}
		pre = root
		root = root.Right
	}
	p1.Val, p2.Val = p2.Val, p1.Val
}

///
func recoverTree1(root *TreeNode) {
	var p1, p2, pre *TreeNode
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != nil && root.Val < pre.Val {
			p1 = root
			if p2 == nil {
				p2 = pre
			}
		}
		pre = root
		dfs(root.Right)
	}
	dfs(root)
	p1.Val, p2.Val = p2.Val, p1.Val
}

// @lc code=end
