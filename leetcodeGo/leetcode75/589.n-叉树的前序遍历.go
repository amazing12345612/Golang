/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
package leetcode75

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var dfs func(root *Node, r []int)
	r := []int{}
	dfs = func(root *Node, r []int) {
		if root != nil {
			r = append(r, root.Val)
		}
		for _, v := range root.Children {
			if v != nil {
				dfs(v, r)
			}
		}
		return
	}
	dfs(root, r)
	return r
}

// @lc code=end
