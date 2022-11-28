/*
 * @lc app=leetcode.cn id=856 lang=golang
 *
 * [856] 括号的分数
 */

// @lc code=start
package tree

func ScoreOfParentheses(s string) int {
	stake := []uint32{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stake = append(stake, uint32(s[i]+'0'))
		} else {
			var sum uint32
			sum = 0
			for stake[len(stake)-1] != '('+'0' {
				sum += stake[len(stake)-1]
				stake = stake[:len(stake)-1]
			}
			stake = stake[:len(stake)-1]
			if sum == 0 {
				stake = append(stake, 1)
			} else {
				stake = append(stake, 2*sum)
			}
		}
	}
	r := 0
	for i := 0; i < len(stake); i++ {
		r += int(stake[i])
	}
	return r
}

// @lc code=end
