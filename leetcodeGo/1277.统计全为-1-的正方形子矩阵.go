/*
 * @lc app=leetcode.cn id=1277 lang=golang
 *
 * [1277] 统计全为 1 的正方形子矩阵
 */

// @lc code=start
package main

func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] == 1 {
					dp[i][j] = 1
					count += dp[i][j]
				}
				continue
			}
			if matrix[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
				count += dp[i][j]
			}
		}
	}
	return count
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end
