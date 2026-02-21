//go:build ignore

// Solution for Exercise 04 — Student Grade Leaderboard
// Run with: go run solution.go
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name   string
	Grades []float64
}

type Rankable interface {
	Score() float64
	GetName() string
}

func (s Student) Score() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0.0
	for _, g := range s.Grades {
		sum += g
	}
	return sum / float64(len(s.Grades))
}

func (s Student) GetName() string {
	return s.Name
}

func leaderboard(students []Rankable) {
	sorted := make([]Rankable, len(students))
	copy(sorted, students)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Score() > sorted[j].Score()
	})
	fmt.Println("=== Leaderboard ===")
	for i, s := range sorted {
		fmt.Printf("  %d. %-10s — %.2f\n", i+1, s.GetName(), s.Score())
	}
}

func main() {
	students := []Rankable{
		Student{Name: "Alice", Grades: []float64{88, 95, 91}},
		Student{Name: "Bob", Grades: []float64{72, 68, 75}},
		Student{Name: "Carol", Grades: []float64{95, 98, 100}},
		Student{Name: "Dave", Grades: []float64{60, 55, 70}},
	}

	leaderboard(students)

	fmt.Println()
	for _, s := range students {
		fmt.Printf("%s score: %.2f\n", s.GetName(), s.Score())
	}
}
