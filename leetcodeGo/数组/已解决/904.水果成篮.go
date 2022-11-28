/*
 * @lc app=leetcode.cn id=904 lang=golang
 *
 * [904] 水果成篮
 */

// @lc code=start
package main

import "fmt"

func totalFruit(fruits []int) int {
	m := map[int]int{}
	l := 0
	r := 0
	ma := 0
	for ; l <= r && r < len(fruits); r++ {
		m[fruits[r]]++
		for len(m) > 2 {
			t := fruits[l]
			m[t]--
			l++
			if m[t] == 0 {
				delete(m, t)
			}
		}
		ma = max(ma, r-l+1)
	}
	return ma
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	fmt.Println(totalFruit([]int{1, 2, 1}))
}

// @lc code=end
