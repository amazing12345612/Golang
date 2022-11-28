/*
 * @lc app=leetcode.cn id=1640 lang=golang
 *
 * [1640] 能否连接形成数组
 */

// @lc code=start
package leetcode75

func CanFormArray(arr []int, pieces [][]int) bool {
	m := make(map[int]int)
	for i, v := range arr {
		m[v] = i + 1
	}
	for _, v := range pieces {
		if m[v[0]] <= 0 {
			return false
		} else {
			t := m[v[0]] - 1
			if len(arr)-t < len(v) {
				return false
			}
			for j := 0; j < len(v) && t < len(arr); j++ {
				if arr[t] != v[j] {
					return false
				}
				t++

			}
		}
	}
	return true
}

// @lc code=end
