/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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

func BuildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	n := len(postorder)
	root := new(TreeNode)
	root.Val = postorder[n-1]
	if n == 1 {
		return root
	}
	//i设为分割点
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[n-1] {
			break
		}
	}
	inleft := inorder[:i]
	postleft := postorder[:i]
	inright := inorder[i+1:]
	postright := postorder[i : n-1]
	root.Left = BuildTree(inleft, postleft)
	root.Right = BuildTree(inright, postright)
	return root

}

// @lc code=end
