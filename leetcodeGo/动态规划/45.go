package dynamicProgramming

/*
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。

*/

func jump(nums []int) int {
	maxPath := 0
	count := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+i > maxPath {
			maxPath = nums[i] + i
		}
		if end == i {
			end = maxPath
			count++
		}
	}
	return count

}
