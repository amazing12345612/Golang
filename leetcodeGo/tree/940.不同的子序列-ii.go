/*
 * @lc app=leetcode.cn id=940 lang=golang
 *
 * [940] 不同的子序列 II
 */

// @lc code=start
package tree

func distinctSubseqII(s string) int {
	n := len(s)
	dp := make([][26]int64, n)
	const mod int64 = 1e9 + 7
	dp[0][s[0]-'a'] = 1
	for i := 1; i < n; i++ {

		for j, v := range dp[i-1] {
			dp[i][s[i]-'a'] += v
			if int(s[i]-'a') != j {
				dp[i][j] = dp[i-1][j]
			}
		}
		dp[i][s[i]-'a'] = 1 + int64(dp[i][s[i]-'a']%mod)
	}
	count := int64(0)
	for _, v := range dp[n-1] {
		count += v
	}
	return int(count)

}

// @lc code=end
