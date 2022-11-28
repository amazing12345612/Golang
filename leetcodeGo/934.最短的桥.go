/*
 * @lc app=leetcode.cn id=934 lang=golang
 *
 * [934] 最短的桥
 */

// @lc code=start
package main

import "fmt"

type pair struct {
	x int
	y int
}

func shortestBridge(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	island := []pair{}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
			return
		}
		grid[i][j] = 2
		p := pair{x: i, y: j}
		island = append(island, p)
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)

	}
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	step := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
			} else {
				continue
			}

			for {
				tmp := island
				island = nil
				for _, v := range tmp {
					for _, d := range dirs {
						x := v.x + d.x
						y := v.y + d.y
						if x >= 0 && x < m && y >= 0 && y < n {
							if grid[x][y] == 1 {
								return step
							}
							if grid[x][y] == 0 {
								grid[x][y] = 2
								island = append(island, pair{x: x, y: y})
							}
						}
					}

				}
				step++
			}

		}
	}
	return step
}
func change(grid [][]int) {
	grid[0][1] = 2
}
func maximumSubarraySum(nums []int, k int) int64 {
	m := map[int]int{}
	if len(nums) < k {
		return 0
	}
	q := []int{}
	sum := 0
	max := 0
	for i := 0; i < k-1; i++ {
		m[nums[i]]++
		q = append(q, nums[i])
		sum += nums[i]
	}
	for i := k - 1; i < len(nums); i++ {
		q = append(q, nums[i])
		m[nums[i]]++
		sum += nums[i]
		if len(m) == k && sum > max {
			max = sum
		}
		sum -= nums[i-k+1]
		q = q[1:]
		m[nums[i-k+1]]--
		if m[nums[i-k+1]] <= 0 {
			delete(m, nums[i-k+1])
		}

	}
	return int64(max)

}
func main() {
	fmt.Println(maximumSubarraySum([]int{2, 3, 3, 2, 1, 3, 4, 5, 5, 5, 1, 2}, 3))
}

// @lc code=end
