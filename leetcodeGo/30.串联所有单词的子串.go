/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
package main

import "fmt"

func findSubstring(s string, words []string) (r []int) {
	n := len(words)
	wordlength := len(words[0])
	m := map[string]int{}
	for _, v := range words {
		m[v]++
	}
	for i := 0; i < len(s)-n*wordlength+1; i++ {
		flag := true
		temp := map[string]int{}
		for k, v := range m {
			temp[k] = v
		}
		for j := 0; j < n; j++ {
			w := s[i+j*wordlength : i+(j+1)*wordlength]
			if temp[w] <= 0 {
				flag = false
				break
			} else {
				temp[w]--
			}
		}

		if flag {
			r = append(r, i)
		}
	}
	return
}
func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}

// @lc code=end
