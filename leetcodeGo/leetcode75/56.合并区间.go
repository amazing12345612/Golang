/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
package leetcode75

import (
	"leetcodeGo/Math"
)

func Merge(intervals [][]int) (r [][]int) {
	QuickSort(intervals, 0, len(intervals)-1)
	for _, v := range intervals {
		if len(r) != 0 && r[len(r)-1][1] >= v[0] {
			r[len(r)-1][1] = Math.Max(v[1], r[len(r)-1][1])
		} else {
			r = append(r, v)
		}
	}

	return
}

func QuickSort(nums [][]int, left, right int) {
	if left >= right {
		return
	}
	j := Jun(nums, left, right)
	QuickSort(nums, left, j-1)
	QuickSort(nums, j+1, right)
}
func Jun(nums [][]int, left, right int) int {
	v := nums[left][0]
	i := left + 1
	j := right
	for {
		for nums[i][0] <= v && i < right {
			i++
		}
		for nums[j][0] >= v && j > left {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]

	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

// @lc code=end
