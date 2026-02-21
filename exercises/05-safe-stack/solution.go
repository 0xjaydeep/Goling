//go:build ignore

// Solution for Exercise 05 — Safe Integer Stack
// Run with: go run solution.go
package main

import (
	"errors"
	"fmt"
)

var ErrEmptyStack = errors.New("stack is empty")

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmptyStack
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmptyStack
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	s := &Stack{}

	fmt.Println("--- Push 10, 20, 30 ---")
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Size:", s.Size())
	fmt.Println("IsEmpty:", s.IsEmpty())

	fmt.Println("\n--- Peek ---")
	if top, err := s.Peek(); err != nil {
		fmt.Println("Peek error:", err)
	} else {
		fmt.Println("Top:", top)
	}
	fmt.Println("Size after Peek:", s.Size())

	fmt.Println("\n--- Pop all ---")
	for !s.IsEmpty() {
		val, _ := s.Pop()
		fmt.Printf("Popped: %d  (size now %d)\n", val, s.Size())
	}

	fmt.Println("\n--- Pop from empty stack ---")
	if _, err := s.Pop(); err != nil {
		fmt.Println("Expected error:", err)
		if errors.Is(err, ErrEmptyStack) {
			fmt.Println("errors.Is confirms: ErrEmptyStack")
		}
	}
}
