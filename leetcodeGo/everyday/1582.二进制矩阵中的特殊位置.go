/*
 * @lc app=leetcode.cn id=1582 lang=golang
 *
 * [1582] 二进制矩阵中的特殊位置
 */

// @lc code=start
package everyday

func NumSpecial(mat [][]int) (r int) {
	m := make([]int, len(mat))
	n := make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			m[i] += mat[i][j]
			n[j] += mat[i][j]
		}
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 && m[i] == 1 && n[j] == 1 {
				r++
			}
		}
	}
	return

}

// @lc code=end
