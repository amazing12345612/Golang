/*
 * @lc app=leetcode.cn id=886 lang=golang
 *
 * [886] 可能的二分法
 */

// @lc code=start
package tree

import "fmt"

func PossibleBipartition(n int, dislikes [][]int) bool {
	m1 := map[int]bool{}
	m2 := map[int]bool{}
	for i := 0; i < len(dislikes); i++ {
		if m1[dislikes[i][0]] && m1[dislikes[i][1]] || m2[dislikes[i][0]] && m2[dislikes[i][1]] {
			return false
		} else if !m1[dislikes[i][0]] && !m1[dislikes[i][1]] && !m2[dislikes[i][1]] && !m2[dislikes[i][0]] {
			m1[dislikes[i][0]] = true
			m2[dislikes[i][1]] = true
		} else if m1[dislikes[i][0]] && !m2[dislikes[i][1]] {
			m2[dislikes[i][1]] = true
		} else {
			m1[dislikes[i][1]] = true
		}
	}
	fmt.Println(3 ^ 1)
	return true
}

// @lc code=end
