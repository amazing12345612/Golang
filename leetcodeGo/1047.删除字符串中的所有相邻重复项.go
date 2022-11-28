/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
package main

import (
	"fmt"
	// "strings"
)

func removeDuplicates(s string) string {
	flag := true
	for flag {
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				s = s[:i] + s[i+2:]
				flag = false
			}
		}
		if !flag {
			flag = true
		} else {
			flag = false
		}
	}
	return s
}
func main() {
	fmt.Println(removeDuplicates("bbacaa"))
}

// @lc code=end
