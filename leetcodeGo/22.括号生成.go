/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
package main

import "fmt"

func generateParenthesis(n int) (result []string) {

	var dfs func(l, r int, path string)
	dfs = func(l, r int, path string) {
		if len(path) == 2*n {
			result = append(result, path)
			return
		}
		if l > 0 {
			dfs(l-1, r, path+"(")
		}
		if l < r {
			dfs(l, r-1, path+")")
		}
	}
	dfs(n, n, "")
	return
}
func main() {
	fmt.Println(generateParenthesis(3))
}

// @lc code=end
