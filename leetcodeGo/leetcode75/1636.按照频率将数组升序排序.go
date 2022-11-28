/*
 * @lc app=leetcode.cn id=1636 lang=golang
 *
 * [1636] 按照频率将数组升序排序
 */

// @lc code=start
package leetcode75

import "sort"

func frequencySort(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if m[i] == m[j] {
			return nums[i] > nums[j]
		}
		return m[i] < m[j]
	})
	return nums
}

// @lc code=end
