package everyday

import "fmt"

func CheckPermutation(s1 string, s2 string) bool {
	m := make(map[byte]int, 0)
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
	}
	fmt.Println(m)
	for i := 0; i < len(s2); i++ {
		if _, ok := m[s2[i]]; ok {
			m[s2[i]]--
			if m[s2[i]] <= 0 {
				delete(m, s2[i])
			}
		} else {
			return false
		}
	}
	return true
}
