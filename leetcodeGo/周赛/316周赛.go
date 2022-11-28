package main

import (
	"fmt"
	"math"
	"strconv"
)

func haveConflict(event1 []string, event2 []string) bool {
	a := tran(event1[0])
	b := tran(event1[1])
	c := tran(event2[0])
	d := tran(event2[1])
	if c >= a && c <= b || d >= a && d <= b || c <= a && d >= b || c >= a && d <= b {
		return true
	} else {
		return false
	}
}
func tran(t string) int {
	a := t[:1]
	b := t[2:]
	c, _ := strconv.Atoi(a)
	d, _ := strconv.Atoi(b)
	return c*60 + d
}
func subarrayGCD(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%k == 0 {
			if nums[i] == k {
				count++
				for j := i + 1; j < len(nums); j++ {
					if nums[j]%k == 0 {
						count++
					} else {
						break
					}
				}
			} else {
				t := nums[i]
				for j := i + 1; j < len(nums); j++ {
					t = gcd(t, nums[j])
					if k < t {
						continue
					} else if t == k {
						count++
					}
				}
			}
		}
	}
	return count
}
func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
func minCost(nums []int, cost []int) int64 {
	min := math.MaxInt64
	max := 0
	var count int64
	var maxc int64
	count = 0
	maxc = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	for i := min; i <= max; i++ {
		for j := 0; j < len(nums); j++ {
			count += int64(abs(nums[j]-i) * cost[j])
		}
		if count > maxc {
			maxc = count
		}
		count = 0
	}
	return int64(maxc)
}
func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
func main() {

	fmt.Println(subarrayGCD([]int{3, 3, 4, 1}, 1))
}
