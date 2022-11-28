/*
 * @lc app=leetcode.cn id=927 lang=golang
 *
 * [927] 三等分
 */

// @lc code=start
package tree

func threeEqualParts(arr []int) []int {
	count := 0
	record := make([]int, len(arr)+1)
	for i, v := range arr {
		if v == 1 {
			count = count + 1
			record[count] = i
		}
	}
	if count == 0 {
		return []int{0, 2}
	}
	if count%3 != 0 {
		return []int{-1, -1}
	}
	gap := count / 3
	count0 := len(arr) - record[count] - 1
	if record[gap*2+1]-record[gap*2]-1 < count0 || record[gap+1]-record[gap]-1 < count0 {
		return []int{-1, -1}
	}
	t1 := record[gap] + count0
	t2 := record[2*gap] + count0
	t3 := record[3*gap] + count0
	nums := min(min(t1+1, t2-t1), t3-t2)
	result := []int{t1, t2 + 1}
	for ; nums > 0; nums-- {
		if arr[t1] != arr[t2] || arr[t1] != arr[t3] || arr[t2] != arr[t3] {
			return []int{-1, -1}
		}
		t1--
		t2--
		t3--
	}
	return result
}

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	} else {
// 		return a
// 	}
// }

// @lc code=end
