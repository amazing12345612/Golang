package dynamicProgramming

import (
	"strconv"
)

func NumDecodings(s string) int {

	//     dp[i] = dp[i-1]+1
	//    if(10 * int(s[i-1])+s[])
	//     s[i]

	//dp[i]表示前i-1个字符
	dp := make([]int, len(s)+1)
	dp[0] = 1

	for i := 1; i <= len(s); i++ {

		if s[i-1] != 48 {
			dp[i] += dp[i-1]
		}
		if i > 1 {
			a, _ := strconv.Atoi(string(s[i-2]))
			b, _ := strconv.Atoi(string(s[i-1]))
			if c := 10*a + b; i > 1 && a != 0 && c <= 26 {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[len(s)]

}
