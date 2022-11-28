package everyday

func goodIndices(nums []int, k int) (r []int) {
	n := len(nums)
	t := true
	dp := make([]bool, n)
	for i := k; i < n-k; i++ {
		t = true

		for j := 0; j < k-1; j++ {
			if nums[j] > nums[j+1] {
				t = false
				break
			}
		}
		if t == false {
			continue
		}
		for j := k + 1; j <= i+k-1; j++ {
			if nums[j] < nums[j+1] {
				t = false
				break
			}
		}
		if t == true {
			r = append(r, i)
			dp[i] = true
		}

	}
	return
}
