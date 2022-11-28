/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
package tree

func Exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	dir := []struct{ x, y int }{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	temp := make([][]bool, m)
	for i := 0; i < m; i++ {
		temp[i] = make([]bool, n)
	}
	var dfs func(i, j int, t int) bool
	dfs = func(i, j int, t int) bool {
		if t == len(word) {
			return true
		}
		for _, v := range dir {
			x := i + v.x
			y := j + v.y

			if x >= m || x < 0 || y < 0 || y >= n || board[x][y] != word[t] || temp[x][y] == true {
				continue
			}
			temp[x][y] = true
			if dfs(x, y, t+1) {
				return true
			} else {
				temp[x][y] = false
			}

		}

		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				temp[i][j] = true
				result := dfs(i, j, 1)
				if result == true {
					return result
				}
				temp[i][j] = false
			}
		}
	}
	return false

}

// @lc code=end
