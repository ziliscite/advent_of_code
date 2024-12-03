package main

import (
	"advent_of_code/utils"
	_ "embed"
	"fmt"
)

// mul(X,Y)
// X and Y are each 1-3 digit
// invalid characters should be ignored
// mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.

// x`mul(2,4)`%&mul[3,7]!@^do_not_`mul(5,5)`+mul(32,64]then(`mul(11,8)``mul(8,5)`)
// Only four highlighted are real mul. 161 (2*4 + 5*5 + 11*8 + 8*5).

//go:embed input.txt
var input string

// -> If m, push. If u and Prev() == m, push, else clear. If l and Prev() == u, push, else clear. Etc, etc.
// if all is right and then ')', return the valid mul and clear

func main() {
	// list to store valid mul(X,Y) answer
	valid := make([]int, 0)

	mul := utils.Stack{}
	num1 := utils.Nums{}
	num2 := utils.Nums{}

	// iterate through each characters
	for i := range input {
		switch r := input[i]; r {
		case 'm':
			if len(mul.Items) == 0 {
				mul.Push("m")
			} else {
				Clear(&mul, &num1, &num2)
			}
		case 'u':
			if mul.Prev() == "m" {
				mul.Push("u")
			} else {
				Clear(&mul, &num1, &num2)
			}
		case 'l':
			if mul.Prev() == "u" {
				mul.Push("l")
			} else {
				Clear(&mul, &num1, &num2)
			}
		case '(':
			if mul.Prev() == "l" {
				mul.Push("(")
			} else {
				Clear(&mul, &num1, &num2)
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// if prev is '(', means that it is for num1, else num2
			if mul.Prev() == "(" && num1.Len() < 3 {
				num1.Push(string(r))
			} else if mul.Prev() == "," && num2.Len() < 3 {
				num2.Push(string(r))
			} else {
				Clear(&mul, &num1, &num2)
			}
		case ',':
			if mul.Prev() == "(" && num1.Len() != 0 {
				mul.Push(",")
			} else {
				Clear(&mul, &num1, &num2)
			}
		case ')':
			if mul.Prev() == "," && num1.Len() > 0 && num2.Len() > 0 {
				n1, _ := num1.Return()
				n2, _ := num2.Return()

				val := n1 * n2
				valid = append(valid, val)

				Clear(&mul, &num1, &num2)
			} else {
				Clear(&mul, &num1, &num2)
			}
		default:
			Clear(&mul, &num1, &num2)
		}
	}

	sum := 0
	for _, p := range valid {
		sum += p
	}

	fmt.Println(sum)
}

func Clear(s *utils.Stack, n1 *utils.Nums, n2 *utils.Nums) {
	s.Clear()
	n1.Clear()
	n2.Clear()
}
