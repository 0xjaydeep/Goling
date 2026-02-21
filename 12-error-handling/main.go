// Demonstrates: errors.New, sentinel errors, custom error type,
// fmt.Errorf with %w (wrapping), errors.Is, errors.As,
// and the idiomatic "if err != nil" pattern.
package main

import (
	"errors"
	"fmt"
)

// --- Sentinel error: a package-level error value to compare against ---
var ErrDivByZero = errors.New("division by zero")

// --- Custom error type: implements the error interface ---
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field %q: %s", e.Field, e.Message)
}

// --- Function returning a sentinel error ---
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

// --- Function returning a custom error type ---
func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically large"}
	}
	return nil
}

// --- Error wrapping: add context while preserving the original error ---
func processAge(age int) error {
	if err := validateAge(age); err != nil {
		return fmt.Errorf("processAge(%d): %w", age, err) // %w wraps, not just formats
	}
	return nil
}

func main() {
	// --- Basic if err != nil pattern ---
	fmt.Println("--- errors.New & if err != nil ---")
	if result, err := divide(10, 3); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	if _, err := divide(5, 0); err != nil {
		fmt.Println("Caught:", err)
	}

	// --- errors.Is: compare through the error chain ---
	fmt.Println("\n--- errors.Is ---")
	_, err := divide(1, 0)
	if errors.Is(err, ErrDivByZero) {
		fmt.Println("errors.Is: confirmed ErrDivByZero")
	}

	// errors.Is unwraps automatically — works even through fmt.Errorf %w chains
	wrapped := fmt.Errorf("outer context: %w", ErrDivByZero)
	fmt.Println("errors.Is on wrapped:", errors.Is(wrapped, ErrDivByZero))

	// --- Custom error type ---
	fmt.Println("\n--- Custom Error Type ---")
	if err := validateAge(-5); err != nil {
		fmt.Println("validateAge(-5):", err)
	}
	if err := validateAge(200); err != nil {
		fmt.Println("validateAge(200):", err)
	}

	// --- errors.As: extract the concrete error type from the chain ---
	fmt.Println("\n--- errors.As ---")
	err2 := processAge(-5) // wraps a *ValidationError inside an extra message
	fmt.Println("processAge error:", err2)

	var ve *ValidationError
	if errors.As(err2, &ve) {
		fmt.Printf("errors.As extracted → field=%q  msg=%q\n", ve.Field, ve.Message)
	}

	// --- Happy path ---
	fmt.Println("\n--- Success Path ---")
	if err := validateAge(25); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age 25 is valid!")
	}
}
