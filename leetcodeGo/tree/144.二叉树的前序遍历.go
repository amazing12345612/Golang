/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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

func PreorderTraversal(root *TreeNode) (result []int) {

	var Dfs func(root *TreeNode)
	Dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		Dfs(root.Left)
		Dfs(root.Right)
	}
	Dfs(root)
	return result
}

func preorderTraversal2(root *TreeNode) (result []int) {
	a := []*TreeNode{}
	a = append(a, root)
	for len(a) > 0 {
		t := a[len(a)-1]
		result = append(result, t.Val)
		a = a[:len(a)-1]
		if t.Right != nil {
			a = append(a, t.Right)
		}
		if t.Left != nil {
			a = append(a, t.Left)
		}
	}
	return result

}

// @lc code=end
