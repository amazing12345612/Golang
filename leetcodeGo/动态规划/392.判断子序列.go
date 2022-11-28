/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
package dynamicProgramming

import "fmt"

func IsSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if j == len(t) && i != len(s) {
		return false
	}
	fmt.Println(i, j)
	return true
}

// @lc code=end
