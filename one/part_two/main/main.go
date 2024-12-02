package main

import (
	"advent_of_code/utils"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	l, r := utils.ParseInputToTwoSlice(input)

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

	fmt.Println(sum)
}
