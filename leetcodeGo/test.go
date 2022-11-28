package main

import (
	"fmt"
)

// "fmt"
// "leetcodeGo/Math"
// "leetcodeGo/tree"
// // "leetcodeGo/leetcode75"
// "leetcodeGo/everyday"
// "reflect"
// "strconv"

func main() {
	// fmt.Println(dynamicProgramming.GenerateParenthesis(3))
	// fmt.Println(dynamicProgramming.Coin(27))
	// fmt.Println(dynamicProgramming.Uniquepath(3, 4))

	// // 64
	// fmt.Println(dynamicProgramming.MinPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	// fmt.Println(dynamicProgramming.MinPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))

	//91
	// fmt.Println(dynamicProgramming.NumDecodings("06"))
	// fmt.Println(dynamicProgramming.NumDecodings("12"))
	// fmt.Println(dynamicProgramming.IsInterleave1("aabcc", "dbbca", "aadbbcbcac"))
	// fmt.Println(dynamicProgramming.Generate(5))

	//121
	// fmt.Println(dynamicProgramming.MaxProfit([]int{1, 2, 3, 4, 5, 3, 5}))

	//139
	// fmt.Println(dynamicProgramming.WordBreak("applepenapple", []string{"apple", "pen"}))

	//152
	// fmt.Println(dynamicProgramming.MaxProduct1([]int{2, 3, -2, 4}))

	//198
	// fmt.Println(dynamicProgramming.Rob([]int{1, 2, 3, 4, 1}))

	//254
	// fmt.Println(dynamicProgramming.NthUglyNumber(10))

	//119
	// fmt.Println(dynamicProgramming.GetRow(3))

	//279
	// fmt.Println(dynamicProgramming.NumSquares(13))

	//300
	// fmt.Println(dynamicProgramming.LengthOfLIS([]int{0, 1, 0, 3, 2, 3}))

	//322
	// fmt.Println(dynamicProgramming.CoinChange([]int{2, 5, 7}, 27))

	//343整数拆分
	// fmt.Println(dynamicProgramming.IntegerBreak(10))

	//1413
	// fmt.Println(everyday.MinStartValue([]int{-3, 2, -3, 4, 2}))

	//377
	// fmt.Println(dynamicProgramming.CombinationSum1_([]int{1, 2, 3}, 4))

	//403
	// fmt.Println(dynamicProgramming.CanCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))

	// //494
	// fmt.Println(dynamicProgramming.FindTargetSumWays([]int{1, 1, 1, 1, 1}, 3))

	// //392
	// fmt.Println(dynamicProgramming.IsSubsequence("axb", "ahbgdc"))

	// //209
	// fmt.Println(dynamicProgramming.MinSubArrayLen1(15, []int{1, 2, 3, 4, 5}))

	// //1219
	// b := [][]int{{0, 0, 0, 3, 0, 0, 17, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 8, 5, 15, 0, 0}, {0, 0, 0, 0, 0, 17, 0, 17, 0, 0, 12, 0}, {0, 0, 0, 4, 0, 0, 20, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 7, 13, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 20}, {2, 0, 0, 4, 0, 0, 0, 16, 0, 4, 0, 0}, {0, 0, 11, 0, 0, 10, 0, 0, 5, 0, 0, 18}, {2, 0, 12, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	// fmt.Println(back.GetMaximumGold1(b))

	//21
	// fmt.Println(21)x`x`
	// p := leetcode75.CreateList([]int{1, 2, 3})
	// q := leetcode75.CreateList([]int{1, 4, 5})
	// p1 := leetcode75.MergeTwoLists(p, q)
	// leetcode75.ShowList(p1)

	// //206
	// p := leetcode75.CreateList([]int{1, 2, 3, 4, 5})
	// q := leetcode75.ReverseList1(p)
	// leetcode75.ShowList(q)
	//299
	// fmt.Println(leetcode75.GetHint("1907", "7910"))
	//793
	// fmt.Println(everyday.ZreimageSizeFZF(6))
	//946
	// fmt.Println(everyday.ValidateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))

	//646 最长数对列

	//56
	// fmt.Println(leetcode75.Merge([][]int{{2, 3}, {1, 2}, {4, 5}, {3, 7}}))
	//698
	// fmt.Println(leetcode75.CanPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))

	//1640 能否形成数组
	// fmt.Println(leetcode75.CanFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}))
	// fmt.Println(leetcode75.CanFormArray([]int{49, 18, 16}, [][]int{{16, 18, 49}}))

	//18四数之和
	// fmt.Println(leetcode75.FourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	//707设计链表
	// head := leetcode75.Constructor()
	// head.AddAtHead(1)
	// head.AddAtTail(3)
	// head.AddAtIndex(1, 2)
	// head.Get(1)
	// head.Show()
	// head.DeleteAtIndex(1)
	// head.Show()

	//1652
	// fmt.Println(leetcode75.Decrypt([]int{2, 4, 9, 3}, -2))
	//106从中序与后序遍历序列构造二叉树
	// t := tree.BuildTree1([]int{8, 4, 15, 12, 7}, []int{8, 15, 7, 12, 4})
	// fmt.Println(tree.PreorderTraversal(t))

	//788 旋转数字
	// fmt.Println(tree.RotatedDigits(10))
	// fmt.Println(everyday.CheckPermutation("abb", "aab"))

	//207
	// fmt.Println(tree.CanFinish(2, [][]int{{0, 1}}))
	//143
	// head := &everyday.ListNode{Val: 1}
	// next := &everyday.ListNode{Val: 2}
	// head.Next = next
	//811子域名访问技术
	// fmt.Println(everyday.SubdomainVisits([]string{"100 abc.baidu.com", "900 baidu.com"}))

	//79单词搜索
	// a := []string{"ABC", "SFC", "ADE"}
	// b := [][]byte{}
	// for _, v := range a {
	// 	b = append(b, []byte(v))
	// }
	// fmt.Println(tree.Exist(b, "ABC"))
	// fmt.Println(tree.ScoreOfParentheses("(())"))
	// fmt.Println(tree.ScoreOfParentheses("((((((())))()())))"))

	//1790
	// fmt.Println(tree.AreAlmostEqual("bank", "kanb"))

	//769
	// fmt.Println(tree.MaxChunksToSorted([]int{1, 0, 2, 3, 5, 4, 6}))

	//886可能的二分法

}
