// Exercise 04 — Student Grade Leaderboard
//
// Problem:
//   Build a student leaderboard using interfaces.
//   Students are ranked by average grade, highest first.
//
// Types/functions to implement:
//   type Student struct { Name string; Grades []float64 }
//   type Rankable interface { Score() float64; GetName() string }
//   (s Student) Score() float64   — average of s.Grades, 0 if empty
//   (s Student) GetName() string  — returns s.Name
//   leaderboard(students []Rankable) — prints ranked list
//
// Hints:
//   sort.Slice(students, fn) — sort by Score() descending
//   fmt.Printf("  %d. %-10s — %.2f\n", rank, name, score)
//
// Run your solution:  go run main.go
// See the answer:     go run solution.go
package main

import "fmt"

// Student holds a name and a slice of exam grades.
type Student struct {
	Name   string
	Grades []float64
}

// Rankable is satisfied by any type that exposes a numeric score and a name.
type Rankable interface {
	Score() float64
	GetName() string
}

// Score returns the average of s.Grades.
// Returns 0.0 if Grades is empty.
func (s Student) Score() float64 {
	// TODO: guard against len(s.Grades) == 0, return 0
	// TODO: sum all grades
	// TODO: return sum / float64(len(s.Grades))
	return 0
}

// GetName returns s.Name.
func (s Student) GetName() string {
	// TODO: return s.Name
	return ""
}

// leaderboard prints students ranked by Score() from highest to lowest.
func leaderboard(students []Rankable) {
	// TODO: make a copy of the slice so you don't mutate the original
	// TODO: sort.Slice(copy, func(i, j int) bool { return copy[i].Score() > copy[j].Score() })
	// TODO: print "=== Leaderboard ==="
	// TODO: for i, s := range sorted { fmt.Printf("  %d. %-10s — %.2f\n", i+1, s.GetName(), s.Score()) }
}

func main() {
	students := []Rankable{
		Student{Name: "Alice", Grades: []float64{88, 95, 91}},
		Student{Name: "Bob", Grades: []float64{72, 68, 75}},
		Student{Name: "Carol", Grades: []float64{95, 98, 100}},
		Student{Name: "Dave", Grades: []float64{60, 55, 70}},
	}

	leaderboard(students)

	// Spot-check individual scores
	fmt.Println()
	for _, s := range students {
		fmt.Printf("%s score: %.2f\n", s.GetName(), s.Score())
	}
}
