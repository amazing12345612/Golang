package main

import (
	"fmt"
	"strconv"
)

func edgeScore(edges []int) int {
	m := make([]int, len(edges))
	temp := 0
	p := 10000
	for i := 0; i < len(edges); i++ {
		m[edges[i]] += i
		if m[edges[i]] > temp {
			temp = m[edges[i]]
			p = edges[i]
		} else if m[edges[i]] == temp && edges[i] < p {
			p = edges[i]
		}
	}
	fmt.Println(m)
	return p

}

func smallestNumber(pattern string) string {
	// s := make([]int, len(pattern)+2)
	s := []int{}
	var result string
	var dfs func(n int)
	count := make(map[int]bool)
	q := 0
	dfs = func(n int) {

		if n == len(pattern)+1 || q != 0 {
			if q == 0 {
				temp := Tstring(s)
				result = temp
				q++
			}
			return
		}
		for i := 1; i <= 9; i++ {
			if count[i] == true {
				continue
			}
			if n >= 1 {
				t := s[n-1]
				if string(pattern[n-1]) == "I" && i <= t {
					return
				}
				if string(pattern[n-1]) == "D" && i >= t {
					return
				}

			}
			count[i] = true
			s = append(s, i)
			fmt.Println(s)
			dfs(n + 1)
			count[i] = false
			s = s[:len(s)-1]

		}

	}
	dfs(0)
	return result
}
func Tstring(s []int) string {
	r := ""
	for _, v := range s {
		temp := strconv.FormatInt(int64(v), 10)
		r = r + temp
	}
	return r
}

func test() {
	result := [][]int{}
	temp := []int{}
	var dfs func(n int)
	t := 0
	dfs = func(n int) {
		if len(temp) == 3 {
			if t == 0 {
				result = append(result, temp)
			}
			t = t + 1
			fmt.Println(temp)
			fmt.Println(result)

			return
		}
		for i := 0; i < 3; i++ {
			temp = append(temp, i)
			dfs(n + 1)
			temp = temp[:len(temp)-1]
		}

	}
	dfs(0)
}
func test1() {
	var baseArr = []int{1, 2, 3}
	var part1 = baseArr[:1] // （1）1
	var part2 = baseArr[1:] // 2,3
	var result [][]int
	for i := 0; i < len(part2); i++ {
		tempArr := append(part1, part2[i]) // （2）
		result = append(result, tempArr)   // （3）
	}

	fmt.Println(result)
}

func main() {
	// fmt.Println(edgeScore([]int{1, 0, 0, 0, 0, 7, 7, 5}))
	// fmt.Println(edgeScore([]int{2, 0, 0, 2}))
	// fmt.Println(smallestNumber("IIID"))
	test()

}
