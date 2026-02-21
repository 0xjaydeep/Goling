//go:build ignore

// Solution for Exercise 01 — Calculator
// Run with: go run solution.go
package main

import "fmt"

func calculate(a float64, op string, b float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator %q", op)
	}
}

func main() {
	type testCase struct {
		a, b float64
		op   string
	}
	cases := []testCase{
		{10, 3, "+"},
		{10, 3, "-"},
		{10, 3, "*"},
		{10, 3, "/"},
		{10, 0, "/"},
		{10, 3, "^"},
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
