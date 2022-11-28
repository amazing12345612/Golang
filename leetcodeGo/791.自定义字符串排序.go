/*
 * @lc app=leetcode.cn id=791 lang=golang
 *
 * [791] 自定义字符串排序
 */

// @lc code=start
package main

import "fmt"

func customSortString(order string, s string) string {
	m := map[rune]int{}
	r := [26]string{}

	for i, v := range order {
		m[v] = i
	}
	temp := ""
	result := ""
	for _, v := range s {
		if _, ok := m[v]; ok {
			r[m[v]] += string(v)
		} else {
			temp += string(v)
		}

	}
	for _, v := range r {
		result += v
	}
	return result + temp
}
func main() {
	customSortString("cba", "cebbabeba")
}

// @lc code=end
