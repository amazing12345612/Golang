/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
package leetcode75

import (
	"sort"
)

func combinationSum(candidates []int, target int) (r [][]int) {
	sort.Ints(candidates)
	nums := []int{}
	var dfs func(nums []int, n int, pre int)
	dfs = func(nums []int, n int, pre int) {
		if n > target {
			return
		}
		if n == target {
			temp := make([]int, len(nums))
			copy(temp, nums)
			r = append(r, temp)
			return
		}

		for i := pre; i < len(candidates); i++ {
			if n+candidates[i] > target {
				return
			}

			nums = append(nums, candidates[i])
			dfs(nums, n+candidates[i], i)
			nums = nums[:len(nums)-1]

		}

	}
	dfs(nums, 0, 0)
	return
}

// @lc code=end
