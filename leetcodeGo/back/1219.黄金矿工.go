/*
 * @lc app=leetcode.cn id=1219 lang=golang
 *
 * [1219] 黄金矿工
 */

// @lc code=start
package back

func GetMaximumGold(grid [][]int) (ans int) {
	var dfs func(i, j int, count int)
	m := len(grid)
	n := len(grid[0])
	var target = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dfs = func(i, j int, count int) {
		t := grid[i][j]
		count += t
		grid[i][j] = 0
		if count > ans {
			ans = count
		}
		for _, v := range target {
			ni := i + v.x
			nj := j + v.y
			if ni >= 0 && nj >= 0 && ni <= m-1 && nj <= n-1 && grid[ni][nj] != 0 {
				dfs(ni, nj, count)
			}
		}
		grid[i][j] = t

	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				dfs(i, j, 0)
			}
		}
	}
	return ans
}

func GetMaximumGold1(grid [][]int) (ans int) {
	var dfs func(i, j int, count int)
	m := len(grid)
	n := len(grid[0])
	dfs = func(i, j int, count int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
			if count > ans {
				ans = count
			}
			return
		}
		t := grid[i][j]
		count += t
		grid[i][j] = 0
		dfs(i-1, j, count)
		dfs(i+1, j, count)
		dfs(i, j-1, count)
		dfs(i, j+1, count)
		count -= t
		grid[i][j] = t

	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				dfs(i, j, 0)
			}
		}
	}
	return ans
}

// @lc code=end
