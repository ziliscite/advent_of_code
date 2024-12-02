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
	d := utils.ParseDiff(l, r)

	sum := 0
	for _, p := range d {
		sum += p
	}

	fmt.Println(sum)
}
