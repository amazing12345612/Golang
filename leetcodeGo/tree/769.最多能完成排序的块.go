/*
 * @lc app=leetcode.cn id=769 lang=golang
 *
 * [769] 最多能完成排序的块
 */

// @lc code=start
package tree

import (
	"sort"
)

func MaxChunksToSorted(arr []int) int {
	back := make([]int, len(arr))
	sum := make([]int, len(arr))
	copy(back, arr)
	copy(sum, arr)
	sort.Ints(sum)
	count := 0
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
		sum[i] += sum[i-1]
	}

	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] == sum[i] {
			count++

		}
	}
	return count
}

// @lc code=end
