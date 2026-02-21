// Exercise 01 — Calculator
//
// Problem:
//   Build a calculator that supports +, -, *, / operations.
//   It must return an error for unknown operators and division by zero.
//
// Functions to implement:
//   calculate(a float64, op string, b float64) (float64, error)
//
// Run your solution:  go run main.go
// See the answer:     go run solution.go
package main

import "fmt"

// calculate performs the arithmetic operation described by op on a and b.
// Supported operators: "+", "-", "*", "/"
// Returns an error if op is unknown or if dividing by zero.
func calculate(a float64, op string, b float64) (float64, error) {
	// TODO: use a switch on op
	// TODO: for "/", check if b == 0 and return fmt.Errorf("division by zero")
	// TODO: for an unknown operator, return fmt.Errorf("unknown operator %q", op)
	return 0, fmt.Errorf("not implemented")
}

func main() {
	type testCase struct {
		a, b float64
		op   string
	}
	cases := []testCase{
		{10, 3, "+"},  // expect 13.0000
		{10, 3, "-"},  // expect  7.0000
		{10, 3, "*"},  // expect 30.0000
		{10, 3, "/"},  // expect  3.3333
		{10, 0, "/"},  // expect error: division by zero
		{10, 3, "^"},  // expect error: unknown operator
	}

	for _, tc := range cases {
		result, err := calculate(tc.a, tc.op, tc.b)
		if err != nil {
			fmt.Printf("calculate(%.0f, %q, %.0f) → ERROR: %v\n", tc.a, tc.op, tc.b, err)
		} else {
			fmt.Printf("calculate(%.0f, %q, %.0f) → %.4f\n", tc.a, tc.op, tc.b, result)
		}
	}
}
