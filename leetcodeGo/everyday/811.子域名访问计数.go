/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
package everyday

import (
	"strconv"
	"strings"
)

func SubdomainVisits(cpdomains []string) (s []string) {
	r := make(map[string]int)
	for _, v := range cpdomains {
		t := strings.Split(v, " ")
		a := strings.Split(t[1], ".")
		num, err := strconv.Atoi(t[0])
		if err != nil {
			return
		}
		r[t[1]] += num
		if len(a) == 3 {
			m := a[1] + "." + a[2]
			n := a[2]
			r[m] += num
			r[n] += num
		} else {
			r[a[1]] += num
		}
	}
	for i, v := range r {
		t := strconv.Itoa(v) + " " + i
		s = append(s, t)

	}

	return
}

// @lc code=end
