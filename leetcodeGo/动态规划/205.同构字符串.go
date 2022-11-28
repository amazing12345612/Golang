/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
package dynamicProgramming

func isIsomorphic(s string, t string) bool {
	a, b := make(map[byte]byte), make(map[byte]byte)
	for i := range s {
		x, y := s[i], t[i]
		if a[x] > 0 && a[x] != y || b[y] > 0 && b[y] != x {
			return false
		}
		a[x] = y
		b[y] = x
	}
	return true
}

// @lc code=end
