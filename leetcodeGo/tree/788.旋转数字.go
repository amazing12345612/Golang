/*
 * @lc app=leetcode.cn id=788 lang=golang
 *
 * [788] 旋转数字
 */

// @lc code=start
package tree

import "fmt"

func RotatedDigits(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		if check(i) {
			fmt.Println(i)
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}
func check(n int) bool {
	r := n
	count := 0
	m := 1
	for n > 0 {
		t := n % 10
		switch t {
		case 1, 8, 0:
			count += t * m
		case 3, 4, 7:
			return false
		case 2:
			count += 5 * m
		case 5:
			count += 2 * m
		case 6:
			count += 9 * m
		case 9:
			count += 6 * 9
		default:
			return false
		}
		m = m * 10
		n = n / 10
	}
	// fmt.Println(count)
	if count != r {
		return true
	} else {
		return false
	}

}

// @lc code=end
