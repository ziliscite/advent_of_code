package test

import (
	"advent_of_code/one/internal"
	_ "embed"
	"sort"
	"testing"
)

//go:embed input.txt
var input string

var left = []int{3, 4, 2, 1, 3, 3}
var right = []int{4, 3, 5, 3, 9, 3}

func TestParseInput(t *testing.T) {
	l, r := internal.ParseInput(input)

	sort.Ints(left)
	sort.Ints(right)

	for i, p := range l {
		if p != left[i] {
			t.Errorf("Expected %d, got %d", left[i], p)
		}
	}

	for i, p := range r {
		if p != right[i] {
			t.Errorf("Expected %d, got %d", right[i], p)
		}
	}
}

func TestParseDiff(t *testing.T) {
	l, r := internal.ParseInput(input)
	d := internal.ParseDiff(l, r)

	for i, p := range d {
		v := l[i] - r[i]
		if v < 0 {
			v = v * -1
		}

		if p != v {
			t.Errorf("Expected %d, got %d", v, p)
		}
	}
}
