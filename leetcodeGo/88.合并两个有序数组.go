/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := len(nums1) - 1
	a := m - 1
	b := n - 1
	for ; i >= 0; i-- {
		if nums1[a] > nums2[b] {
			nums1[i] = nums1[a]
			a--
		} else {
			nums1[i] = nums1[b]
			b--
		}
		if b < 0 {
			return
		}
	}
	return
}
func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	merge(a, 3, b, 3)
	fmt.Println(a)
}

// @lc code=end
