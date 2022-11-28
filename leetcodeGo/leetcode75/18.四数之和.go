/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
package leetcode75

func FourSum(nums []int, target int) [][]int {
	r := [][]int{}
	m := map[int]bool{}
	temp := []int{}
	var dfs func(sum int, n int, idx int)
	if len(nums) < 4 {
		return r
	}
	dfs = func(sum int, n int, idx int) {
		if n == 4 && sum == target {
			new := make([]int, len(temp))
			copy(new, temp)
			r = append(r, new)
		}
		for i := idx; i < len(nums); i++ {
			if m[i] || i > 0 && len(temp) > 0 && nums[i] == temp[len(temp)-1] {
				continue
			}
			// if i == len(nums)-1 && len(temp) != 4 {
			// 	return
			// }
			m[i] = true
			temp = append(temp, nums[i])
			dfs(sum+nums[i], n+1, idx+1)
			temp = temp[:len(temp)-1]
			m[i] = false
		}
	}
	dfs(0, 0, 0)
	return r
}

// @lc code=end
