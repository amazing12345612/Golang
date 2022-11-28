/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
package main

func longestValidParentheses(s string) int {
	stake := []int{}
	stake = append(stake, -1)
	m := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stake = append(stake, i)
		} else {
			stake = stake[:len(stake)-1]
			if len(stake) == 0 {
				stake = append(stake, i)
			} else {
				m = max(m, i-stake[len(stake)-1])
			}

		}
	}
	return m
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
