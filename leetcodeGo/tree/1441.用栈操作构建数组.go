/*
 * @lc app=leetcode.cn id=1441 lang=golang
 *
 * [1441] 用栈操作构建数组
 */

// @lc code=start
package tree

func buildArray(target []int, n int) []string {
	a := []string{"Push", "Pop"}
	s := []string{}
	j := 0
	for i := 1; i <= n && j < len(target); i++ {
		s = append(s, a[0])
		if i != target[j] {
			s = append(s, a[1])
		} else {
			j++
		}
	}
	return s
}

// @lc code=end
