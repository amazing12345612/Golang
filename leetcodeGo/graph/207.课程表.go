/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
package tree

import "fmt"

func CanFinish(numCourses int, prerequisites [][]int) bool {
	//入度表
	indegree := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	for _, v := range prerequisites {
		indegree[v[0]]++
		adjacency[v[1]] = append(adjacency[v[1]], v[0])
	}
	fmt.Println(indegree, adjacency)
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	fmt.Println(queue)
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		numCourses--
		for _, v := range adjacency[i] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return numCourses == 0

}

// @lc code=end
