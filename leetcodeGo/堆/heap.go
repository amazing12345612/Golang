package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[:n-1]
	*h = old[0 : n-1]
	return x
}

func howMany(S string, k int) int {
	// write code here
	m := map[byte]int{}
	count := 0
	for i := 0; i < len(S); i++ {
		m[S[i]]++
	}
	for _, v := range m {
		if v >= k {
			count++
		}
	}
	return count
}
func minCnt(s string) int {
	// write code here
	count := 0
	if len(s) == 1 {
		return 0
	}
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			count++
		}
	}
	return count
}
func main() {

	// a := heap.Remove(h, 0)
	nums := []int{1, 1, 2, 3, 4, 5, 6}
	res := []float64{}
	big := IntHeap{}
	small := IntHeap{}
	heap.Init(&small)
	heap.Init(&big)
	res[0] = float64(nums[0])
	if nums[0] == nums[1] {
		heap.Push(&big, nums[0])
		heap.Push(&small, -nums[1])
		res[1] = float64(nums[0])
	} else if nums[0] > nums[1] {
		heap.Push(&small, -nums[0])
		heap.Push(&big, nums[1])
		res[1] = (float64(nums[0]) + float64(nums[1])) / 2.0
	} else {
		heap.Push(&small, -nums[1])
		heap.Push(&big, nums[0])
		res[1] = (float64(nums[0]) + float64(nums[1])) / 2.0
	}

	for i := 2; i < len(nums); i++ {
		if small.Len() == big.Len() {
			if nums[i] >= -small[0] && nums[i] <= big[0] {
				heap.Push(&small, nums[i])
			} else if nums[i] < -small[0] {
				heap.Push(&small, nums[i])
			} else {
				heap.Push(&big, nums[i])
			}
		} else {

		}

	}

	//fmt.Println()

}
