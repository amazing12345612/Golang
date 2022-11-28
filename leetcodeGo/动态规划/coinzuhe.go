package dynamicProgramming

import ()

// func min(a, b, c int) int {
// 	temp := 0
// 	if a > b {
// 		temp = b
// 	} else {
// 		temp = a
// 	}
// 	if temp > c {
// 		return c
// 	} else {
// 		return temp
// 	}
// }
//用2，5，7面额硬币组合成固定的数值的最少硬币数量
// dp[i]表示组合成I数值所需要的最少硬币数
// dp[i] = min (dp[i-2]+1,dp[i-5]+1,dp[i-7]+1)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Coin(n int) int {
	dp := make([]int, n+1)
	nums := []int{2, 5, 7}
	Maxnum := 100
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = Maxnum
		for _, v := range nums {
			if i-v >= 0 {
				dp[i] = min(dp[i-v]+1, dp[i])
			}
		}
	}
	return dp[27]
}
