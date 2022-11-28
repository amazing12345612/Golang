/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)

}

//确定函数参数
//从根节点向下寻找 找到符合条件的就返回
func hasPathSum2(root *TreeNode, count int) bool {
	//确定返回条件
	if root.Left == nil && root.Right == nil && count == 0 {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return false
	}
	//确定递归逻辑
	if root.Left != nil {
		count -= root.Val
		if hasPathSum2(root.Left, count) {
			return true
		}
		count += root.Val
	}
	if root.Right != nil {
		count -= root.Val
		if hasPathSum2(root.Right, count) {
			return true
		}
		count += root.Val
	}
	return false

}

// @lc code=end
