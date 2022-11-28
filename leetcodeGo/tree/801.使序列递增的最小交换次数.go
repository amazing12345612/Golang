/*
 * @lc app=leetcode.cn id=801 lang=golang
 *
 * [801] 使序列递增的最小交换次数
 */

// @lc code=start
package tree

func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = 1
	for i := 1; i < n+1; i++ {

		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] && nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			dp[i][0] = min(dp[i-1][0], dp[i-1][1])
			dp[i][1] = min(dp[i-1][1], dp[i-1][0]) + 1
		} else if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			dp[i][1] = dp[i-1][0] + 1
			dp[i][0] = dp[i-1][1]
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + 1
		}

	}
	return min(dp[n-1][0], dp[n-1][1])
}

// @lc code=end
