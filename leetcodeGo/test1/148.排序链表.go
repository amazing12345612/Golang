/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type HeapNode []*ListNode

func (h HeapNode) Len() int           { return len(h) }
func (h HeapNode) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h HeapNode) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *HeapNode) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))

}
func (h *HeapNode) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x

}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	h := &HeapNode{}
	heap.Init(h)
	p := head
	for p != nil {
		fmt.Println(p.Val)
		heap.Push(h, p)
		p = p.Next
	}
	l := new(ListNode)
	p1 := l
	for h.Len() > 0 {
		n := heap.Pop(h).(*ListNode)
		n.Next = nil
		l.Next = n
		l = l.Next
	}
	return p1.Next
}

//
func find(nums []int, k int) {
	m := map[int]int{}
	count := 0
	for _, v := range nums {
		if m[v] > 0 {
			count += m[v]
		} else {
			m[k-v]++
		}
	}
	fmt.Println(count)

}

func main() {
	// a := []int{4, 2, 3, 1}
	// head1 := new(ListNode)
	// p1 := head1
	// for _, v := range a {
	// 	temp := &ListNode{Val: v}
	// 	p1.Next = temp
	// 	p1 = p1.Next
	// }
	// result := sortList(head1.Next)
	// for result != nil {

	// 	fmt.Print(result.Val)
	// 	result = result.Next
	// }
	find([]int{1, 2, 3, 4, 5}, 6)
}

// @lc code=end
