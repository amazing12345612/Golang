/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			dp[i] = dp[i-1] + prices[i] - prices[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(prices)-1]

}

// @lc code=end
