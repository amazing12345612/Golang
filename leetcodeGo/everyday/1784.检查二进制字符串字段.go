/*
 * @lc app=leetcode.cn id=1784 lang=golang
 *
 * [1784] 检查二进制字符串字段
 */

// @lc code=start
package everyday

func checkOnesSegment(s string) bool {
	i := 0
	for ; i < len(s); i++ {
		if s[i] == '0' {
			break
		}
	}
	if i == len(s)-1 {
		return true
	}
	i = i + 1
	for ; i < len(s); i++ {
		if s[i] == '1' {
			return false
		}
	}
	return true

}

// @lc code=end
