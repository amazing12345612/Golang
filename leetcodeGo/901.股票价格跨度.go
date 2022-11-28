/*
 * @lc app=leetcode.cn id=901 lang=golang
 *
 * [901] 股票价格跨度
 */

// @lc code=start
package main

import (
	"fmt"
)

type Price struct {
	x   int
	day int
}
type StockSpanner struct {
	Stake []Price
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	t := Price{}
	t.x = price
	t.day = 1
	for len(this.Stake) > 0 && price >= this.Stake[len(this.Stake)-1].x {
		a := this.Stake[len(this.Stake)-1]
		this.Stake = this.Stake[:len(this.Stake)-1]
		t.day += a.day
	}
	this.Stake = append(this.Stake, t)
	return t.day
}
func main() {
	a := Constructor()
	b := []int{100, 80, 60, 70, 60, 75, 85}
	for _, v := range b {
		fmt.Println(a.Next(v))
	}
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end
