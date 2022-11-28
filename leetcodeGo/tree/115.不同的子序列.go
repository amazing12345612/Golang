/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
package tree

func numDistinct(s string, t string) int {
	m := len(t)
	n := len(s)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m+1; i++ {
		dp2 := make([]int, n+1)
		for j := 1; j < n+1; j++ {
			if s[i-1] != t[j-1] {
				dp2[j] = dp2[j-1]
			} else {
				dp2[j] = dp2[j-1] + dp[j-1]
			}
		}
		dp = dp2
	}
	return dp[n]

}

// @lc code=end
