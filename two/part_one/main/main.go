package main

import (
	"advent_of_code/utils"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	l := utils.ParseInputTo2DSlice(input)

	c := 0
	for i, p := range l {
		tn := len(p) - 1

		ni := make([]int, 0)
		for j := range tn {
			dif := l[i][j] - l[i][j+1]
			ni = append(ni, dif)
		}

		y := true
		prev := ni[0]
		for k := range tn {
			if ni[k] > 0 && prev < 0 {
				y = false
				continue
			}

			if ni[k] < 0 && prev > 0 {
				y = false
				continue
			}

			if ni[k] == 0 {
				y = false
				continue
			}

			if ni[k] > 0 {
				if ni[k] > 3 || ni[k] < 0 {
					y = false
					continue
				}
			} else {
				if ni[k] < -3 || ni[k] > 0 {
					y = false
					continue
				}
			}

			prev = ni[k]
		}

		if y {
			c++
		}
	}

	fmt.Println(c)
}
