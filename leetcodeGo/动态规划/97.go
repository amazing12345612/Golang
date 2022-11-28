package dynamicProgramming

func IsInterleave(s1 string, s2 string, s3 string) bool {

	m := 0
	for i, j := 0, 0; m < len(s1)+len(s2); {
		if i < len(s1) && s1[i] == s3[m] {
			i++
			m = i + j
			continue
		}
		if j < len(s2) && s2[j] == s3[m] {
			j++
			m = i + j
			i, j = j, i
			s1, s2 = s2, s1
			continue
		}
		return false

	}
	if len(s1)+len(s2) == len(s3) {
		return true
	}
	return false

}

func IsInterleave1(s1 string, s2 string, s3 string) bool {

	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i >= 1 && dp[i-1][j] == true {
				dp[i][j] = (s1[i-1] == s3[i+j-1])
			}
			if j >= 1 && dp[i][j-1] == true {
				dp[i][j] = (s2[j-1] == s3[i+j-1])
			}

		}
	}
	return dp[m][n]

}
