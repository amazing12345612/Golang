package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] > h[j] } //最小堆
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {}
func (h *IntHeap) Pop() interface{}   { return nil }

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

}
