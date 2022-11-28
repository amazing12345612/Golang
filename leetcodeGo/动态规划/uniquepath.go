package dynamicProgramming

//给定mxn的矩阵，🤖️从左上角走到右下角有多少种走法，🤖️每次可以向右或向下走

func Uniquepath(m, n int) int {

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]

			}

		}
	}
	return dp[m-1][n-1]

}
