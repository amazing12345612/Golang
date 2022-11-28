/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
package dynamicProgramming

func FindTargetSumWays(nums []int, target int) int {

	var dfs func(int)
	sum := 0
	count := 0
	dfs = func(n int) {
		if n == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		for i := -1; i < 2; i = i + 2 {
			sum += (i * nums[n])
			dfs(n + 1)
			sum -= (i * nums[n])
		}
	}
	dfs(0)
	return count
}

func FindTargetSumWays1(nums []int, target int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, target+1)
	}
	for i := 0; i < len(nums); i++ {

	}
}

// @lc code=end
