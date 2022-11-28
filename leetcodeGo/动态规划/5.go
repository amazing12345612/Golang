package dynamicProgramming

func longestpalindromic(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := ""
	for i := len(s); i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
func longestpalindromic1(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = middlefloat(s, i, i, res)
		res = middlefloat(s, i, i+1, res)
	}
	return res
}
func middlefloat(s string, i, j int, res string) string {
	temp := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		temp = s[i : j+1]
		i--
		j++
	}
	if len(temp) > len(res) {
		return temp
	}
	return res
}
