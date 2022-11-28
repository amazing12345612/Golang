/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start
package leetcode75

import (
	"sort"
)

func CanPartitionKSubsets(nums []int, k int) bool {
	c := 0
	//记录nums中数字使用情况,int表示下标
	count := map[int]bool{}
	//求和 算均值
	for _, v := range nums {
		c += v
	}
	//若无法整除k说明不存在分割可能
	if c%k != 0 {
		return false
	}
	//t子集和
	t := c / k
	//n子集数 sum 当前子集和
	sort.Ints(nums)
	var dfs func(sum int, n int, idx int) bool
	dfs = func(sum int, n int, idx int) bool {
		if n == k {
			return true
		}
		if sum == t {
			return dfs(0, n+1, idx-1)
		}
		if idx == -1 {
			return false
		}
		for i := idx; i >= 0; i-- {
			if count[i] || sum+nums[i] > t {
				continue
			}
			count[i] = true
			if dfs(sum+nums[i], n, idx-1) {
				return true
			}
			count[i] = false
			if sum == 0 {
				return false
			}
		}
		return false
	}
	return dfs(0, 0, len(nums)-1)
}

// @lc code=end
