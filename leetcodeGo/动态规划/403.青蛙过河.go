/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */

// @lc code=start
package dynamicProgramming

import "fmt"

func CanCross(stones []int) bool {
	target := stones[len(stones)-1]

	a := make(map[int]bool, len(stones)-1)
	for _, v := range stones {
		a[v] = true
	}
	fmt.Println(a)
	return test(1, target, 1, a)

}

func test(a, target, c int, m map[int]bool) bool {
	if a == target {
		return true
	} else if m[a] != true {
		return false
	} else {
		return test(a+c, target, c, m) || test(a+c-1, target, c-1, m) || test(a+c+1, target, c+1, m)
	}
}

// @lc code=end
