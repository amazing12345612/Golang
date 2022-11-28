/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
package dynamicProgramming

import (
	"fmt"
	"leetcodeGo/Math"
)

func Rob213(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return Math.Max(nums[0], nums[1])
	} else {

		return Math.Max(Rob1(nums[1:]), Rob1(nums[:n-1]))
	}
}
func Rob1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = Math.Max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = Math.Max(dp[i-2]+nums[i], dp[i-1])

	}
	fmt.Println(dp)
	return dp[n-1]
}

// @lc code=end
