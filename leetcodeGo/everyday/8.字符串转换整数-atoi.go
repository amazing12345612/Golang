/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
package everyday

func myAtoi(s string) int {
	r := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '.' || s[i]-'A' >= 0 && s[i]-'A' < 26 || s[i]-'a' >= 0 && s[i]-'a' < 26 && len(r) == 0 {
			return 0
		}
		r = r + ""
	}

	return 1
}

// @lc code=end
