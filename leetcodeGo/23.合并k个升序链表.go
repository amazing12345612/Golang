/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
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

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var dfs func(root *ListNode, lists []*ListNode) *ListNode
	dfs = func(root *ListNode, lists []*ListNode) *ListNode {
		if len(lists) == 0 {
			return root
		}
		root2 := lists[0]
		return dfs(Merge(root, root2), lists[1:])
	}
	return dfs(lists[0], lists[1:])

}
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var dfs func(lists []*ListNode, l, r int) *ListNode
	dfs = func(lists []*ListNode, l, r int) *ListNode {
		if l == r {
			return lists[l]
		}
		if l > r {
			return nil
		}
		mid := l + (l-r)/2
		return Merge(dfs(lists, l, mid), dfs(lists, mid+1, r))
	}
	return dfs(lists, 0, len(lists)-1)

}

func Merge(root1 *ListNode, root2 *ListNode) *ListNode {
	p1 := root1
	p2 := root2
	pn := new(ListNode)
	t := pn
	for p1 != nil || p2 != nil {
		if p1 == nil {
			pn.Next = p2
			break
		}
		if p2 == nil {
			pn.Next = p1
			break
		}
		if p1.Val < p2.Val {
			pn.Next = p1
			pn = pn.Next
			p1 = p1.Next
		} else {
			pn.Next = p2
			pn = pn.Next
			p2 = p2.Next
		}

	}
	return t.Next
}
func main() {
	t := new(ListNode)
	t.Val = 1
	b := new(ListNode)
	b.Val = 3
	t.Next = b
	t1 := new(ListNode)
	t1.Val = 2
	b1 := new(ListNode)
	b1.Val = 3
	t1.Next = b1
	new := merge(t, t1)
	for new != nil {
		fmt.Println(new.Val)
		new = new.Next
	}

}

// @lc code=end
