/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
package main

import "fmt"

func trap(height []int) int {
	result := 0
	for i := 1; i <= MaxHeight(height); i++ {
		flag := false
		temp := 0
		ans := 0
		for _, v := range height {

			if flag && v < i {
				temp++
			}
			if v >= i {
				ans += temp
				temp = 0
				flag = true
			}
		}
		result += ans
	}

	return result
}
func MaxHeight(height []int) (r int) {
	for _, v := range height {
		if v > r {
			r = v
		}
	}
	return
}
func trap1(height []int) int {
	s := []int{}
	r := 0
	cur := 0
	for cur < len(height) {
		for len(s) > 0 && height[cur] > height[s[len(s)-1]] {
			t := height[s[len(s)-1]]
			s = s[:len(s)-1]
			if len(s) == 0 {
				break
			}
			d := cur - s[len(s)-1] - 1
			min := min(height[cur], height[s[len(s)-1]])
			r += d * (min - t)

		}
		s = append(s, cur)
		cur++

	}
	return r
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func main() {
	h1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// h2 := []int{4, 2, 0, 3, 2, 5}
	// fmt.Println(MaxHeight(h1))
	fmt.Println(trap1(h1))
}

// @lc code=end
