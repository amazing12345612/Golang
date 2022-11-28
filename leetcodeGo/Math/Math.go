package Math

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func Min3(a, b, c int) int {
	t := Min(a, b)
	return Min(t, c)
}

func Max3(a, b, c int) int {
	t := Max(a, b)
	return Max(t, c)

}
