/*
 * @lc app=leetcode.cn id=907 lang=golang
 *
 * [907] 子数组的最小值之和
 */

// @lc code=start
package main

import "fmt"

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	s := []int{-1}
	s2 := []int{n}
	for i, v := range arr {
		for len(s) > 1 && v <= arr[s[len(s)-1]] {
			s = s[:len(s)-1]
		}
		left[i] = s[len(s)-1]
		s = append(s, i)
	}
	for i := n - 1; i >= 0; i-- {
		for len(s2) > 1 && arr[i] < arr[s2[len(s2)-1]] {
			s2 = s2[:len(s2)-1]
		}
		right[i] = s2[len(s2)-1]
		s2 = append(s2, i)
	}
	sum := 0
	for i, v := range arr {
		sum += v * (i - left[i]) * (right[i] - i) % (1e9 + 7)
	}
	fmt.Println(left, right)
	return sum
}
func main() {
	fmt.Println(sumSubarrayMins([]int{1, 4, 2, 3, 1}))
}

// @lc code=end
