/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
package dynamicProgramming

import "fmt"

func IntegerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	dp[1] = 1
	for i := 3; i < n+1; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(max(dp[i], dp[i-j]*j), (i-j)*j)
		}
	}
	fmt.Println(dp)
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
