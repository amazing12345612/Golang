/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
package main

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	r := n + 1
	s := make([]int, r)
	s[0] = 0
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + nums[i]
	}
	i := 0
	for j := 0; j < n; j++ {
		for s[j]-s[i] >= target && i < j {
			r = min(j-i, r)
			i++
		}
	}
	if r == n+1 {
		return 0
	}
	return r
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end
