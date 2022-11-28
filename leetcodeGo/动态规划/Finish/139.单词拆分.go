/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start

func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s)+1)
	Minlen := len(wordDict[0])
	for _, v := range wordDict {
		if len(v) < Minlen {
			Minlen = len(v)
		}
	}
	if len(s) < Minlen {
		return false
	}
	dp[0] = true
	for i := 0; i < len(s)+1; i++ {
		if i < Minlen {
			continue
		}
		for _, v := range wordDict {
			if i-len(v) >= 0 {
				dp[i] = dp[i] || (dp[i-len(v)] && (s[:i-len(v)]+v == s[:i]))
			}

		}
	}
	fmt.Println(dp)
	// dp[i] = dp[i-1] && wordDict[0]+ s[:i-len(wordDict[0])] == s[:i]
	return dp[len(s)]
}

// @lc code=end
