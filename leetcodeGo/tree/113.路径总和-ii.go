/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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

func pathSum(root *TreeNode, targetSum int) [][]int {
	temp := []int{}
	result := [][]int{}
	var pathSum1 func(root *TreeNode, tatargetSum int)
	pathSum1 = func(root *TreeNode, tatargetSum int) {
		if root.Left == nil && root.Right == nil && targetSum == 0 {
			new := make([]int, len(temp))
			copy(new, temp)
			result = append(result, new)
			return
		}
		if root.Left == nil && root.Right == nil {
			return
		}
		if root.Left != nil {
			temp = append(temp, root.Left.Val)
			pathSum1(root.Left, targetSum-root.Left.Val)
			temp = temp[:len(temp)-1]
		}
		if root.Right != nil {
			temp = append(temp, root.Right.Val)
			pathSum1(root.Right, targetSum-root.Right.Val)
			temp = temp[:len(temp)-1]
		}
	}

	if root == nil {
		return [][]int{}
	} else {
		temp = append(temp, root.Val)
		pathSum1(root, targetSum-root.Val)
	}
	return result
}

// @lc code=end
