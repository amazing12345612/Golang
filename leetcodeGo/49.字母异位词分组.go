/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, v := range strs {
		t := []byte(v)
		sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
		s := string(t)
		m[s] = append(m[s], v)
	}
	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
func main() {
	a := "abc"
	t := []int{}
	for _, v := range a {
		t = append(t, int(v))
	}
	sort.Ints(t)

}

// @lc code=end
