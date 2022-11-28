package leetcode75

import "fmt"

func CreateList(a []int) *ListNode {
	head := new(ListNode)
	end := head
	head.Val = 0
	for i := 0; i < len(a); i++ {
		temp := new(ListNode)
		temp.Val = a[i]
		end.Next = temp
		end = end.Next
	}
	return head.Next
}

func ShowList(head *ListNode) {
	r := []int{}
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	fmt.Println(r)
}
