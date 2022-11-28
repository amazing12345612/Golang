/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
package dynamicProgramming

import (
	"fmt"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = 10000
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i-j*j], dp[i])
		}
		dp[i] = dp[i] + 1
	}
	fmt.Println(dp)
	return dp[n]
}

// @lc code=end
