package test

import (
	"advent_of_code/utils"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestSolution(t *testing.T) {
	l := utils.ParseInputTo2DSlice(input)

	c := 0
	for _, p := range l {
		g := false

		for n, _ := range p {
			np := make([]int, len(p)-1)
			copy(np[:n], p[:n])
			copy(np[n:], p[n+1:])

			tn := len(np) - 1

			ni := make([]int, 0)
			for o := range tn {
				dif := np[o] - np[o+1]
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
				g = true
			}
		}

		if g {
			c++
		}
	}

	fmt.Println(c)
}
