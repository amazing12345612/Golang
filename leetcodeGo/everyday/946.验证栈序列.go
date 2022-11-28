/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */

// @lc code=start
package everyday

func ValidateStackSequences(pushed []int, popped []int) bool {
	r := []int{}
	j := 0
	for i := 0; i < len(pushed); i++ {
		r = append(r, pushed[i])

		for len(r) > 0 && r[len(r)-1] == popped[j] {
			r = r[:len(r)-1]
			j++
		}
	}
	if len(r) == 0 {
		return true
	}
	return false
}

// @lc code=end
