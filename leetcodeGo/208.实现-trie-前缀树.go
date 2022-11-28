/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this.children[0]
	for _, v := range word {
		v -= 'a'
		if this.children[v] != nil {
			node = this.children[v]
		} else {
			this.children[v] = &Trie{}
			node = this.children[v]
		}
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := &Trie{}
	for _, v := range word {
		if this.children[v] != nil {
			node = this.children[v]
		} else {
			return false
		}
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := &Trie{}
	for _, v := range prefix {
		if this.children[v] != nil {
			node = this.children[v]
		} else {
			return false
		}
	}
	fmt.Println(node)
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
