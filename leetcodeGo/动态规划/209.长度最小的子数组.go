/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
package dynamicProgramming

import (
	"fmt"
	"sort"
)

func MinSubArrayLen(target int, nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	min := 1000000
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target && i < min {
			min = i + 1
		}
		for j := i; j >= 0; j-- {
			if nums[i]-nums[j] >= target && i-j < min {
				min = i - j
			}
		}
	}
	fmt.Println(nums)
	if min != 1000000 {
		return min
	}
	return 0
}

func MinSubArrayLen1(target int, nums []int) int {
	n := len(nums)
	sum := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	Min := 1000000
	for i := 1; i < len(nums)+1; i++ {
		m := target + sum[i-1]
		j := sort.Search(n+1, func(i int) bool { return sum[i] >= m })
		if j > 0 && j != n+1 {
			Min = min(Min, j-i+1)
		}
	}

	fmt.Println(nums)

	if Min != 1000000 {
		return Min
	}
	return 0
}

// @lc code=end
