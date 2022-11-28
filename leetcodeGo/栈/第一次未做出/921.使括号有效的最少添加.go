/*
 * @lc app=leetcode.cn id=921 lang=golang
 *
 * [921] 使括号有效的最少添加
 */

// @lc code=start
package tree

func minAddToMakeValid(s string) int {
	count1 := 0
	count2 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count1++
		} else {
			if count1 != 0 {
				count1--
			} else {
				count2++
			}
		}
	}
	return count2

}

// @lc code=end
