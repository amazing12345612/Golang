/*
 * @lc app=leetcode.cn id=862 lang=golang
 *
 * [862] 和至少为 K 的最短子数组
 */

// @lc code=start
package main

import "fmt"

func shortestSubarray(nums []int, k int) (r int) {
	s := make([]int, len(nums)+1)
	s[0] = 0
	for i := 0; i < len(nums); i++ {
		s[i+1] = s[i] + nums[i]
	}
	q := []int{}
	r = len(nums) + 1
	for i, v := range s {
		for len(q) > 0 && v-s[q[0]] >= k {
			r = min(r, i-q[0])
			q = q[1:]
		}
		for len(q) > 0 && v <= s[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func main() {
	fmt.Println(shortestSubarray([]int{1, 2, 3}, 1))
}

// @lc code=end
