/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
package dynamicProgramming

import (
	"fmt"
	"leetcodeGo/Math"
)

func NthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	if n == 1 {
		return 1
	}
	p1, p2, p3 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = Math.Min3(dp[p1]*2, dp[p2]*3, dp[p3]*5)
		if dp[i] == dp[p1]*2 {
			p1 = p1 + 1
		}
		if dp[i] == dp[p2]*3 {
			p2 = p2 + 1
		}
		if dp[i] == dp[p3]*5 {
			p3 = p3 + 1
		}
	}
	fmt.Println(dp)
	return dp[n-1]

}

// @lc code=end
