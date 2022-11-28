/*
 * @lc app=leetcode.cn id=870 lang=golang
 *
 * [870] 优势洗牌
 */

// @lc code=start
package tree

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	r := make([]int, len(nums1))
	sort.Ints(nums1)
	ids := make([]int, len(nums2))
	for i := range nums2 {
		ids[i] = i
	}
	sort.Slice(ids, func(i int, j int) bool { return nums2[ids[i]] < nums2[ids[i]] })
	left := 0
	right := len(nums2) - 1
	for _, v := range nums1 {
		if v > nums2[ids[left]] {
			r[ids[left]] = v
			left++
		} else {
			r[ids[right]] = v
			right--
		}
		if left > right {
			break
		}
	}
	return r
}

// @lc code=end
