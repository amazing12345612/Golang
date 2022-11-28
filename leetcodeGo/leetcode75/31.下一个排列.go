/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
package leetcode75

import "sort"

func nextPermutation(nums []int) {
	j := -1
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			j = i
			t := len(nums) - 1
			for ; t >= i; t-- {
				if nums[t] > nums[i-1] {
					break
				}
			}
			nums[t], nums[i-1] = nums[i-1], nums[t]
			break

		}
	}
	if j == -1 {
		sort.Ints(nums)
		return
	}
	sort.Ints(nums[j:])

}

// @lc code=end
