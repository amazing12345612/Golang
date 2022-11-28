/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
package leetcode75

import "sort"

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	r := 1
	for _, v := range nums {
		if v > r {
			break
		} else if v == r {
			r++
		}
	}
	return r
}
func firstMissingPositive1(nums []int) int {
	dp := make([]int, len(nums)-1)
	dp[0] = Judge(nums[0])
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] == dp[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums)-1]
}
func Judge(n int) int {
	if n == 1 {
		return 2
	} else {
		return 1
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end
