/*
 * @lc app=leetcode.cn id=1413 lang=golang
 *
 * [1413] 逐步求和得到正数的最小值
 */

// @lc code=start
package everyday

func MinStartValue(nums []int) int {
	sum := 0
	Min := 0
	for _, v := range nums {
		sum += v
		Min = min(sum, Min)
	}
	return 1 - Min
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end
