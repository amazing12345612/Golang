/*
 * @lc app=leetcode.cn id=617 lang=golang
 *
 * [617] 合并二叉树
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

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	n := &TreeNode{}
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 != nil && root2 != nil {
		n.Val = root1.Val + root2.Val
		n.Left = mergeTrees(root1.Left, root2.Left)
		n.Right = mergeTrees(root1.Right, root2.Right)
	} else {
		if root1 != nil {
			n = root1
		} else {
			n = root2
		}
	}
	return n

}

// @lc code=end
