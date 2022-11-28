/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
package dynamicProgramming

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	Map1 := make(map[int]int)
	for i := 0; i < n; i++ {
		r := nums[i] % k
		if v, ok := Map1[r]; ok {
			if i-v >= 2 {
				return true
			}
		} else {
			Map1[r] = i
		}
	}
	return false
}

// @lc code=end
