/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
package main

func isValid(s string) bool {
	stake := []rune{}
	m := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stake = append(stake, v)
		} else {
			top := stake[len(s)-1]
			stake = stake[:len(stake)-1]
			if v != m[top] {
				return false
			}
		}
	}
	return true
}

// @lc code=end
