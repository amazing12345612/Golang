/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
package tree

func NumTrees(n int) int {
	g := make([]int, n+1)
	g[0] = 1
	g[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}
	return g[n]
}

func generate(l, r int) int {
	if l > r {
		return 0
	}
	result := 0
	for i := l; i <= r; i++ {
		m := generate(l, i-1)
		n := generate(i-1, r)
		result = m * n
	}
	return result
}

// @lc code=end
