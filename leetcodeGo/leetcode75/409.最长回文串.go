/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
package leetcode75

func longestPalindrome(s string) int {
	m1 := make(map[byte]bool)
	count := 0
	for i := 0; i < len(s); i++ {
		if !m1[s[i]] {
			m1[s[i]] = true
		} else {
			delete(m1, s[i])
			count += 2
		}
	}
	if len(m1) != 0 {
		count = count + 1
	}
	return count
}

// @lc code=end
