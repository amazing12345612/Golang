package dynamicProgramming

import (
	"leetcodeGo/Math"
)

func CanJump(nums []int) bool {
	maxPath := nums[0]
	for i, v := range nums {
		if maxPath >= len(nums)-1 {
			return true
		}
		if i <= maxPath { //如果位置可以小于最远可到达距离 则表示可到达
			if i+v > maxPath { //更新最远距离
				maxPath = i + v
			}
		}
	}
	return false
}
func CanJump2(nums []int) bool {
	dp := make([]int, len(nums)) //存储I位置能够最远达到的
	dp[0] = nums[0]
	if len(nums) == 1 {
		return true
	}
	if dp[0] == 0 {
		return false
	}
	for i := 1; i < len(nums)-1; i++ { //考虑到倒数第二个下标
		dp[i] = Math.Max(dp[i-1], i+nums[i])
		if dp[i] == i { //如果在最后一个下标前，最远到达就已经是只能是当前I位置，说明无法达到最后一个下标
			return false
		}
	}
	return true
}

func CanJump3(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	temp := nums[0]
	if temp == 0 {
		return false
	}
	for i := 1; i < len(nums)-1; i++ {
		n := Math.Max(temp, i+nums[i])
		if n == i {
			return false
		}
		temp = n
	}
	return true

}
