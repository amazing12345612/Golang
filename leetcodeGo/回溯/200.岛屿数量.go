/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
package main

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	count := 0
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'

		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}

		}

	}
	return count
}

// func main() {
// 	nums := [][]byte{{'1', '1', '1'}}
// 	fmt.Println(numIslands(nums))
// }

// @lc code=end
