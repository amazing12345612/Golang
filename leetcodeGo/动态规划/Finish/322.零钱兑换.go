/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start

var Max uint32 = ^uint32(0) >> 1

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = int(Max)
		for _, v := range coins {
			if i-v >= 0 && dp[i-v] != -1 {
				dp[i] = min(dp[i-v]+1, dp[i])
			}

		}
		if dp[i] == int(Max) {
			dp[i] = -1
		}
	}
	fmt.Println(dp)
	return dp[amount]
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end
