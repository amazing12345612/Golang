/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 */

// @lc code=start
package main

import "fmt"

func letterCasePermutation(s string) (r []string) {
	m := []int{}
	for i, v := range s {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			m = append(m, i)
		}
	}
	temp := []byte(s)
	var dfs func(temp []byte, target, i int)
	dfs = func(temp []byte, target int, i int) {

	}

	return
}
func main() {
	fmt.Println(letterCasePermutation("abc"))
}

// @lc code=end
