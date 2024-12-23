package main

import (
	"advent_of_code/utils"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	// list to store valid mul(X,Y) answer
	valid := make([]int, 0)

	mul := utils.Stack{}
	do := utils.Instruction{
		IsDo: true,
	}

	num1 := utils.Nums{}
	num2 := utils.Nums{}

	// iterate through each characters
	for i := range input {
		switch r := input[i]; r {
		case 'm':
			if len(mul.Items) == 0 && do.IsDo {
				mul.Push("m")
			} else {
				Clear(&mul, &num1, &num2, &do)
			}
		case 'u':
			if mul.Prev() == "m" && do.IsDo {
				mul.Push("u")
			} else {
				Clear(&mul, &num1, &num2, &do)
			}
		case 'l':
			if mul.Prev() == "u" && do.IsDo {
				mul.Push("l")
			} else {
				Clear(&mul, &num1, &num2, &do)
			}
		case '(':
			if mul.Prev() == "l" && do.IsDo {
				mul.Push("(")
			} else if do.Prev() == "o" || do.Prev() == "t" {
				do.Push("(")
			} else {
				Clear(&mul, &num1, &num2, &do)
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// if prev is '(', means that it is for num1, else num2
			if mul.Prev() == "(" && num1.Len() < 3 && do.IsDo {
				num1.Push(string(r))
			} else if mul.Prev() == "," && num2.Len() < 3 && do.IsDo {
				num2.Push(string(r))
			} else {
				Clear(&mul, &num1, &num2, &do)
			}
		case ',':
			if mul.Prev() == "(" && num1.Len() != 0 && do.IsDo {
				mul.Push(",")
			} else {
				Clear(&mul, &num1, &num2, &do)
			}
		case ')':
			if mul.Prev() == "," && num1.Len() > 0 && num2.Len() > 0 && do.IsDo {
				n1, _ := num1.Return()
				n2, _ := num2.Return()

				val := n1 * n2
				valid = append(valid, val)
			} else if do.PrevPrev() == "o" {
				do.IsDo = true
			} else if do.PrevPrev() == "t" {
				do.IsDo = false
			}

			Clear(&mul, &num1, &num2, &do)
		case 'd':
			if len(do.Items) == 0 {
				do.Push("d")
			} else {
				Clear(&mul, &num1, &num2, &do)
			}
		case 'o':
			if do.Prev() == "d" {
				do.Push("o")
			} else {
				Clear(&mul, &num1, &num2, &do)
			}
		case 'n':
			if do.Prev() == "o" {
				do.Push("n")
			} else {
				Clear(&mul, &num1, &num2, &do)
			}
		case '\'':
			if do.Prev() == "n" {
				do.Push("'")
			} else {
				Clear(&mul, &num1, &num2, &do)
			}
		case 't':
			if do.Prev() == "'" {
				do.Push("t")
			} else {
				Clear(&mul, &num1, &num2, &do)
			}
		default:
			Clear(&mul, &num1, &num2, &do)
		}
	}

	sum := 0
	for _, p := range valid {
		sum += p
	}

	fmt.Println(sum)
}

func Clear(s *utils.Stack, n1 *utils.Nums, n2 *utils.Nums, ins *utils.Instruction) {
	s.Clear()
	n1.Clear()
	n2.Clear()
	ins.Clear()
}
