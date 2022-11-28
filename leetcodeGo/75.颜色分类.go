/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
package main

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for i, v := range nums {
		if v == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		} else if v == 2 {
			if nums[right] == 0 {
				nums[right], nums[left] = nums[left], nums[right]
				left++
			} else {
				nums[right], nums[i] = nums[i], nums[right]
				right--
			}

		}
	}
}

// @lc code=end
