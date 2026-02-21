// Exercise 05 — Safe Integer Stack
//
// Problem:
//   Implement a LIFO stack for integers that returns errors instead of
//   panicking when you try to Pop or Peek an empty stack.
//
// Types/functions to implement:
//   type Stack struct { items []int }
//   (s *Stack) Push(val int)
//   (s *Stack) Pop() (int, error)   — ErrEmptyStack if empty
//   (s *Stack) Peek() (int, error)  — ErrEmptyStack if empty, does NOT remove
//   (s *Stack) IsEmpty() bool
//   (s *Stack) Size() int
//
// Run your solution:  go run main.go
// See the answer:     go run solution.go
package main

import (
	"errors"
	"fmt"
)

// ErrEmptyStack is returned by Pop and Peek when the stack has no elements.
var ErrEmptyStack = errors.New("stack is empty")

// Stack is a last-in-first-out (LIFO) collection of integers.
type Stack struct {
	// TODO: add a slice field to store elements
	//   items []int
}

// Push adds val to the top of the stack.
func (s *Stack) Push(val int) {
	// TODO: s.items = append(s.items, val)
}

// Pop removes and returns the top element.
// Returns (0, ErrEmptyStack) if the stack is empty.
func (s *Stack) Pop() (int, error) {
	// TODO: if s.IsEmpty() { return 0, ErrEmptyStack }
	// TODO: top := s.items[len(s.items)-1]
	// TODO: s.items = s.items[:len(s.items)-1]
	// TODO: return top, nil
	return 0, fmt.Errorf("not implemented")
}

// Peek returns the top element without removing it.
// Returns (0, ErrEmptyStack) if the stack is empty.
func (s *Stack) Peek() (int, error) {
	// TODO: if s.IsEmpty() { return 0, ErrEmptyStack }
	// TODO: return s.items[len(s.items)-1], nil
	return 0, fmt.Errorf("not implemented")
}

// IsEmpty reports whether the stack has no elements.
func (s *Stack) IsEmpty() bool {
	// TODO: return len(s.items) == 0
	return true
}

// Size returns the number of elements currently in the stack.
func (s *Stack) Size() int {
	// TODO: return len(s.items)
	return 0
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
	fmt.Println("Size after Peek:", s.Size()) // should still be 3

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
