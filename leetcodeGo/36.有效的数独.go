/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
package main

func isValidSudoku(board [][]byte) bool {
	row := [9][9]int{}
	column := [9][9]int{}
	s := [3][3][9]int{}
	for i, v1 := range board {
		for j, v2 := range v1 {
			if v2 == '.' {
				continue
			}
			row[i][v2]++
			column[j][v2]++
			s[i/3][j/3][v2]++
			if row[i][v2] > 1 || column[j][v2] > 1 || s[i/3][j/3][v2] > 1 {
				return false
			}
		}
	}
	return true
}

// @lc code=end
