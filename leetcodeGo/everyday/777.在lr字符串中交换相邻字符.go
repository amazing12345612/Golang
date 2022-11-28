/*
 * @lc app=leetcode.cn id=777 lang=golang
 *
 * [777] 在LR字符串中交换相邻字符
 */

// @lc code=start
package everyday

func canTransform(start string, end string) bool {
	n := len(start)
	i, j := 0, 0
	for i < n || j < n {
		for i < n && start[i] == 'x' {
			i++
		}
		for j < n && end[j] == 'x' {
			j++
		}
		if i == n || j == n {
			return i == j
		}
		if start[i] != end[j] {
			return false
		}
		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		i++
		j++
	}
	return i == j
}

// @lc code=end
