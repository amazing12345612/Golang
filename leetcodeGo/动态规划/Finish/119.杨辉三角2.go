package dynamicProgramming

func GetRow(rowIndex int) []int {

	dp := make([]int, rowIndex+1)
	dp[0] = 1
	for i := 0; i < rowIndex+1; i++ {
		a := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				a[j] = 1
			} else {
				a[j] = dp[j-1] + dp[j]
			}
		}
		dp = a
		if i == rowIndex {
			return a
		}
	}
	return []int{}
}
