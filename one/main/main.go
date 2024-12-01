package main

import (
	"advent_of_code/one/internal"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	l, r := internal.ParseInput(input)
	d := internal.ParseDiff(l, r)

	sum := 0
	for _, p := range d {
		sum += p
	}

	fmt.Println(sum)
}