/*
 * @lc app=leetcode.cn id=1790 lang=golang
 *
 * [1790] 仅执行一次字符串交换能否使两个字符串相等
 */

// @lc code=start
package main

import "fmt"

func AreAlmostEqual(s1 string, s2 string) bool {
	var m1 byte
	var m2 byte
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
			if count > 2 {
				return false
			}
			if count == 1 {
				m1 = s1[i]
				m2 = s2[i]
				continue
			}
			if s1[i] == m2 && s2[i] == m1 {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
func main() {
	fmt.Println(AreAlmostEqual("123", "123"))
}

// @lc code=end
