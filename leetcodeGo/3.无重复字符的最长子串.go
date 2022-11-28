/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := map[byte]int{}
	j := 0
	max := 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for j < len(s) && m[s[j]] == 0 {
			m[s[j]]++
			j++
		}
		if j-i > max {
			max = j - i
		}
	}
	return max
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

// @lc code=end
