/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
package main

func combinationSum(candidates []int, target int) (result [][]int) {
	temp := []int{}
	var dfs func()
	dfs = func() {
		if sum(temp) == target {
			result = append(result, append([]int{}, temp...))
			return
		}
		for i := 0; i < len(candidates); i++ {
			if sum(temp)+candidates[i] > target {
				continue
			}
			temp = append(temp, candidates[i])
			dfs()
			temp = temp[:len(temp)-1]
		}
	}
	dfs()
	return
}
func sum(nums []int) (r int) {
	for _, v := range nums {
		r += v
	}
	return

}

// @lc code=end
