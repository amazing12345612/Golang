/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 颠倒字符串中的单词
 */

// @lc code=start
package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == ' ' && s[i+1] == ' ' {
			strings.Replace(s, " ", "", i)
		}
	}
	fmt.Println(s)
}

// @lc code=end
