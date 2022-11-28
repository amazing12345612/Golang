package tree

func maxDepth(root *TreeNode) int {
	m := 0
	var dfs func(root *TreeNode, n int)
	dfs = func(root *TreeNode, n int) {
		if root == nil {
			return
		}
		if n > m {
			m = n
		}
		dfs(root.Left, n+1)
		dfs(root.Right, n+1)

	}
	dfs(root, 1)
	return m
}
