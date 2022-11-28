/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
package leetcode75

func setZeroes(matrix [][]int) {
	m := map[int]bool{}
	n := map[int]bool{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				m[i] = true
				n[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if m[i] == true || n[j] == true {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroes1(matrix [][]int) {
	r0 := false
	c0 := false
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			r0 = true
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			c0 = true
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0

			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if r0 == true {
		matrix[0] = make([]int, len(matrix[0]))
	}
	if c0 == true {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end
