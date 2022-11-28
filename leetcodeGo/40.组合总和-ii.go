/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) (result [][]int) {
	sort.Ints(candidates)
	temp := []int{}
	// count := map[int]bool{}
	var dfs func(n int)
	dfs = func(n int) {
		if sum(temp) == target {
			result = append(result, append([]int{}, temp...))
			return
		}
		for i := n; i < len(candidates); i++ {
			if sum(temp)+candidates[i] > target || (i > n && candidates[i] == candidates[i-1]) {
				continue
			}
			// count[i] = true
			temp = append(temp, candidates[i])
			dfs(i + 1)
			// count[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0)
	return
}
func sum(nums []int) (r int) {
	for _, v := range nums {
		r += v
	}
	return

}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func main() {
	a := []int64{}
	a = append(a, 1, 3, 4, 5, 6)

	fmt.Println(cap(a))
}

// @lc code=end
