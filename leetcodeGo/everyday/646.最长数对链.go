/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */

// @lc code=start
package everyday

func findLongestChain(pairs [][]int) int {
	Sort(pairs)
	r := 1
	t := pairs[0][1]
	for _, v := range pairs {

		if v[0] > t {
			r++
			t = v[1]
		}
	}
	return r

}
func Sort(pairs [][]int) [][]int {
	for i := 1; i < len(pairs); i++ {
		for j := i; j-1 >= 0; j-- {
			if pairs[j][1] < pairs[j-1][1] {
				temp := pairs[j]
				pairs[j] = pairs[j-1]
				pairs[j-1] = temp
			}
		}
	}
	return pairs
}

// @lc code=end
