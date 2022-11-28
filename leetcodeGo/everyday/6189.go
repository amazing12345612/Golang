package everyday

func longestSubarray(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	count := 0
	t := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == max {
			count++
		} else {
			t = Max(count, t)
			count = 0
		}

	}
	return t
}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
