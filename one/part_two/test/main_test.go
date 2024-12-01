package test

import (
	"advent_of_code/utils"
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestSolution(t *testing.T) {
	l, r := utils.ParseInput(input)

	sl := make([]int, len(l))
	for i, _ := range l {
		for _, p := range r {
			if l[i] == p {
				sl[i]++
			}
		}
	}

	sum := 0
	for i, p := range sl {
		sum += l[i] * p
	}

	if sum != 31 {
		t.Errorf("Expected 31, got %d", sum)
	}
}
