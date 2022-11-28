/*
 * @lc app=leetcode.cn id=1608 lang=golang
 *
 * [1608] 特殊数组的特征值
 */

// @lc code=start
package everyday

import (
	"sort"
)

func specialArray(nums []int) int {
	m := len(nums)
	sort.Ints(nums)
	for i := 0; i <= m; i++ {
		for i, v := range nums {
			if v >= i && len(nums[i:]) == i {
				return i
			}
		}
	}
	return -1

}

// @lc code=end
