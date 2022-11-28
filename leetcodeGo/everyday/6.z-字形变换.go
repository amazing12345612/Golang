/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
package everyday

func Convert(s string, numRows int) string {
	r := make([]string, numRows)
	flag := -1
	n := 0
	for i := 0; i < len(s); i++ {
		r[n] += string(s[i])
		if n == 0 || n == numRows-1 {
			flag = -flag
		}
		n += flag

	}
	t := ""
	for _, v := range r {
		t += v
	}
	return t
}

// @lc code=end
