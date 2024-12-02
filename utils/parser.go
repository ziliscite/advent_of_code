package utils

import (
	"sort"
	"strconv"
	"strings"
)

// ParseInputToTwoSlice splits the input string into two sorted slices of int
func ParseInputToTwoSlice(input string) ([]int, []int) {
	s := strings.ReplaceAll(input, "\r\n", "\n")
	n := strings.ReplaceAll(s, "\r", "\n")

	d := strings.Split(n, "\n")

	f := make([]string, 0)
	for _, p := range d {
		f = append(f, strings.Split(p, "   ")...)
	}

	l := make([]int, 0)
	r := make([]int, 0)
	for i, p := range f {
		n, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			l = append(l, n)
		} else {
			r = append(r, n)
		}
	}

	sort.Ints(l)
	sort.Ints(r)

	return l, r
}

// ParseInputTo2DSlice parse the input string into a 2d slice of int
func ParseInputTo2DSlice(input string) [][]int {
	s := strings.ReplaceAll(input, "\r\n", "\n")
	n := strings.ReplaceAll(s, "\r", "\n")

	d := strings.Split(n, "\n")

	l := make([][]int, 0)
	for _, p := range d {
		nl := strings.Split(p, " ")

		il := make([]int, len(nl))
		for i, m := range nl {
			n, err := strconv.Atoi(m)
			if err != nil {
				panic(err)
			}

			il[i] = n
		}

		l = append(l, il)
	}

	return l
}
