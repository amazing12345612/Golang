package main

func restoreIpAddresses(s string) []string {
	// write code here
	temp := ""
	ans := []string{}
	var dfs func(n int, pre int)
	dfs = func(n, pre int) {
		if n == 4 && pre == len(s)-1 {
			t := temp
			ans = append(ans, t[:len(t)-1])
			return
		}
		for i := pre + 1; i < len(s); i++ {
			if !check(s[pre:i]) {
				return
			}
			temp += s[pre:i] + "."
			dfs(n+1, i)
			temp = temp[:len(temp)-i+pre-1]

		}

	}
	dfs(0, 0)
	return ans
}

func check(nums string) bool {
	n := len(nums)
	if nums[0] == '0' && len(nums) > 1 {
		return false
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum*10 + int(nums[i]-'0')
	}
	if sum > 255 {
		return false
	} else {
		return true
	}

}
