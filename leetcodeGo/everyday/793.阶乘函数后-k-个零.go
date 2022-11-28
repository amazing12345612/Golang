/*
 * @lc app=leetcode.cn id=793 lang=golang
 *
 * [793] 阶乘函数后 K 个零
 */

// @lc code=start
package everyday

import (
	"sort"
)

func Zeta(n int) (res int) {
	for n > 0 {
		n /= 5
		res += n
	}
	return
}

func nx(k int) int {
	return sort.Search(5*k, func(x int) bool { return Zeta(x) >= k })
}

func ZreimageSizeFZF(k int) int {
	return nx(k)
}

// @lc code=end
