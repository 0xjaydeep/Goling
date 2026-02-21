// Package calculator provides basic arithmetic operations.
//
// Exported identifiers (uppercase first letter) are accessible from other packages.
// Unexported identifiers (lowercase) are private to this package.
package calculator

import "errors"

// unexported helper — only visible inside the calculator package.
func isZero(n float64) bool {
	return n == 0
}

// Add returns the sum of a and b.
func Add(a, b float64) float64 {
	return a + b
}

// Sub returns a minus b.
func Sub(a, b float64) float64 {
	return a - b
}

// Mul returns the product of a and b.
func Mul(a, b float64) float64 {
	return a * b
}

// Div returns a divided by b.
// Returns an error if b is zero.
func Div(a, b float64) (float64, error) {
	if isZero(b) {
		return 0, errors.New("calculator: division by zero")
	}
	return a / b, nil
}
