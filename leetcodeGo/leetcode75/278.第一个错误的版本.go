/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
package leetcode75

func isBadVersion(i int) bool {
	if i > 10 {
		return true
	} else {
		return false
	}
}
func firstBadVersion(n int) int {
	l := 1
	r := n
	for l <= r {
		mid := (r-l)/2 + l
		if isBadVersion(mid) && !isBadVersion(mid-1) && mid > 1 {
			return mid
		} else if !isBadVersion(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return 1
}

// @lc code=end
