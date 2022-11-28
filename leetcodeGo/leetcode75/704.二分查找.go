/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
package leetcode75

func search(nums []int, target int) (a int) {

	var search1 func(nums []int, l, r int) int
	search1 = func(nums []int, l, r int) int {
		mid := (l + r) / 2
		if l > r {
			return -1
		}
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			return search1(nums[:mid], l, mid-1)
		} else {
			return search1(nums[mid+1:], mid+1, l)
		}
	}
	a = search1(nums, 0, len(nums)-1)
	return a
}

// @lc code=end
