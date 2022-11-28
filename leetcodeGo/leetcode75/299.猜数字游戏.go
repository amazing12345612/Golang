/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */

// @lc code=start
package leetcode75

import (
	"strconv"
)

func GetHint(secret string, guess string) string {
	n := make([]int, 10)
	m := make([]int, 10)
	count := 0
	count1 := 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			count++
		} else {
			n[secret[i]-'0']++
			m[guess[i]-'0']++
		}
	}
	for i := 0; i < len(n); i++ {
		a := min(n[i], m[i])
		count1 += a
	}
	d := strconv.Itoa(count)
	e := strconv.Itoa(count1)

	return d + "A" + e + "B"
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end
