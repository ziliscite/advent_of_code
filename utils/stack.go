package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Nums to keep tracks of valid mul
type Nums struct {
	items []string
}

func (s *Nums) Push(data string) {
	s.items = append(s.items, data)
}

func (s *Nums) Len() int {
	return len(s.items)
}

func (s *Nums) Clear() {
	if len(s.items) == 0 {
		return
	}

	s.items = make([]string, 0)
}

func (s *Nums) Return() (int, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("empty")
	}

	str := strings.Join(s.items, "")

	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	s.Clear()

	return num, nil
}

// Stack to keep tracks of valid mul
type Stack struct {
	Items []string
}

func (s *Stack) Push(data string) {
	s.Items = append(s.Items, data)
}

func (s *Stack) Prev() string {
	if len(s.Items) == 0 {
		return ""
	}

	return s.Items[len(s.Items)-1]
}

func (s *Stack) Clear() {
	if len(s.Items) == 0 {
		return
	}

	s.Items = make([]string, 0)
}

func (s *Stack) Return() string {
	if len(s.Items) == 0 {
		return ""
	}

	mul := strings.Join(s.Items, "")
	s.Clear()

	return mul
}

// Instruction to keep tracks of valid instructions
type Instruction struct {
	Items []string
	IsDo  bool
}

func (s *Instruction) Push(data string) {
	s.Items = append(s.Items, data)
}

func (s *Instruction) Prev() string {
	if len(s.Items) == 0 {
		return ""
	}

	return s.Items[len(s.Items)-1]
}

func (s *Instruction) PrevPrev() string {
	if len(s.Items) == 0 {
		return ""
	}
	return s.Items[len(s.Items)-2]
}

func (s *Instruction) Clear() {
	if len(s.Items) == 0 {
		return
	}
	s.Items = make([]string, 0)
}
