/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start

package dynamicProgramming

import (
	"fmt"
	"leetcodeGo/Math"
)

func MaxProduct(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {

		sum := 1
		max := nums[i]
		for j := i; j >= 0; j-- {
			sum *= nums[j]
			if sum > max {
				max = sum
			}
		}
		dp[i] = max

	}
	Maxnum := dp[0]
	for _, v := range dp {
		if v > Maxnum {
			Maxnum = v
		}
	}
	return Maxnum
}

func MaxProduct1(nums []int) int {
	n := len(nums)
	dpmin, dpmax := nums[0], nums[0]
	Maxnum := dpmax
	for i := 1; i < n; i++ {
		m, n := dpmax, dpmin
		dpmin = Math.Min3(nums[i], nums[i]*n, nums[i]*m)
		dpmax = Math.Max3(nums[i], nums[i]*n, nums[i]*m)
		if dpmax > Maxnum {
			Maxnum = dpmax
		}
	}
	fmt.Println(dpmax, dpmin)
	return Maxnum

}

// @lc code=end
