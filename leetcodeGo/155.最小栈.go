/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
package main

import "math"

type MinStack struct {
	s        []int
	minStake []int
}

func Constructor() MinStack {

	return MinStack{minStake: []int{math.MaxInt32}}
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s, val)
	if val < this.minStake[len(this.minStake)-1] {
		this.minStake = append(this.minStake, val)
	} else {
		this.minStake = append(this.minStake, this.minStake[len(minStake)-1])
	}

}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
	this.minStake = this.minStake[:len(this.minStake)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStake[len(this.minStake)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
