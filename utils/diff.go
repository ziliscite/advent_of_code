package utils

// ParseDiff returns the absolute difference between each element of two slices. Input must be a sorted slice
func ParseDiff(l, r []int) []int {
	n := make([]int, 0)

	for i := 0; i < len(l); i++ {
		v := l[i] - r[i]

		if v < 0 {
			n = append(n, v*-1)
		} else {
			n = append(n, v)
		}
	}

	return n
}
