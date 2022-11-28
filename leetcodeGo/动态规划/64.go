package dynamicProgramming

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

*/

//方法一 动态规划
import (
	"leetcodeGo/Math"
)

func MinPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else {
				if i == 0 && j >= 1 {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				} else if j == 0 && i >= 1 {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				} else {
					dp[i][j] = Math.Min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
				}
			}
		}
	}
	return dp[m-1][n-1]

}

//优化
//不需要额外dp数组空间 直接在grid原数组上修改
