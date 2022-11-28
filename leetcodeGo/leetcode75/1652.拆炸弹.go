/*
 * @lc app=leetcode.cn id=1652 lang=golang
 *
 * [1652] 拆炸弹
 */

// @lc code=start
package leetcode75

import "fmt"

func Decrypt(code []int, k int) []int {
	n := len(code)
	t := make([]int, n)
	code = append(append(code, code...), code...)
	fmt.Println(code, n)
	for i := 0; i < n; i++ {
		if k == 0 {
			code[i] = 0
		} else if k > 0 {
			sum := 0
			for j := 1; j <= k; j++ {
				sum += code[n+j+i]
			}
			t[i] = sum
		} else {
			sum := 0
			for j := 1; j <= -k; j++ {
				sum += code[n-j+i]
			}
			t[i] = sum
		}
	}
	return t
}

// @lc code=end
