/*
 * @lc app=leetcode.cn id=1800 lang=golang
 *
 * [1800] 最大升序子数组和
 */

// @lc code=start
package main

func maxAscendingSum(nums []int) int {
	max := 0
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
func main() {
	println(maxAscendingSum([]int{1, 2, 3}))

}

// @lc code=end
