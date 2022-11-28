/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
package main

import "fmt"

func longestPalindrome(s string) string {
	max := ""
	var middle func(i, j int)
	middle = func(i, j int) {
		temp := ""
		for i >= 0 && j < len(s) && s[i] == s[j] {
			temp = s[i : j+1]
			i--
			j++
		}
		if len(temp) > len(max) {
			max = temp
		}
		return
	}
	for i := 0; i < len(s); i++ {
		middle(i, i)
		middle(i, i+1)
	}
	return max
}
func main() {
	fmt.Println(longestPalindrome("babad"))
}

// @lc code=end
