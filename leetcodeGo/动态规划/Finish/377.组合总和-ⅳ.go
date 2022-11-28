/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
package dynamicProgramming

import "fmt"

func CombinationSum4(nums []int, target int) int {
	var dfs func(int)
	sum := 0
	count := 0
	dfs = func(n int) {
		if sum == target {
			count++
			return
		}
		for i := 0; i < len(nums); i++ {
			if sum > target {
				continue
			}
			sum = sum + nums[i]
			dfs(n + 1)
			sum = sum - nums[i]

		}
	}
	dfs(1)
	return count
}

func CombinationSum1_(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i < target+1; i++ {
		for _, v := range nums {
			if i-v >= 0 {
				dp[i] += dp[i-v]
			}
		}
	}
	fmt.Println(dp)
	return dp[target]

}

// @lc code=end
