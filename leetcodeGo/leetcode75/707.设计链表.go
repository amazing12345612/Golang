/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 */

// @lc code=start
package leetcode75

import (
	"fmt"
)

type MyLinkedList struct {
	ListNode *ListNode
	Size     int
}

func Constructor() MyLinkedList {
	Linklist := &MyLinkedList{&ListNode{}, 0}
	return *Linklist
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Size || index < 0 {
		return -1
	}
	p := this.ListNode
	for i := index - 1; i >= 0; i-- {
		p = p.Next
	}
	fmt.Println(p.Val)
	return p.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	index = Max(0, index)
	if index > this.Size {
		return
	}
	this.Size++
	if index == 0 {
		p := this
		temp := &ListNode{val, nil}
		temp.Next = p.ListNode
		p.ListNode = temp
		return
	}
	pre := this.ListNode
	for i := index - 1; i > 0; i-- {
		pre = pre.Next
	}
	temp := &ListNode{val, pre.Next}
	pre.Next = temp
	return

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	this.Size--
	//t 表示要删除的节点的前节点
	if index == 0 {
		this.ListNode = this.ListNode.Next
		return
	}
	t := this.ListNode
	for i := index - 1; i > 0; i-- {
		t = t.Next
	}
	temp := t.Next
	t.Next = temp.Next
	temp.Next = nil
	return

}

func (this *MyLinkedList) Show() {
	r := []int{}
	p := this.ListNode
	for p.Next != nil {
		r = append(r, p.Val)
		p = p.Next
	}
	fmt.Println(r)
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end
