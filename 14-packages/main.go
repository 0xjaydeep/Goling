// Demonstrates: importing a local package, calling exported functions,
// and the exported vs unexported (public vs private) visibility rule.
package main

import (
	"fmt"

	"go-basics-practice/14-packages/calculator"
)

func main() {
	fmt.Println("--- Local Package: calculator ---")
	a, b := 12.0, 4.0

	fmt.Printf("Add(%.0f, %.0f) = %.2f\n", a, b, calculator.Add(a, b))
	fmt.Printf("Sub(%.0f, %.0f) = %.2f\n", a, b, calculator.Sub(a, b))
	fmt.Printf("Mul(%.0f, %.0f) = %.2f\n", a, b, calculator.Mul(a, b))

	if result, err := calculator.Div(a, b); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Div(%.0f, %.0f) = %.2f\n", a, b, result)
	}

	// --- Error path ---
	if _, err := calculator.Div(a, 0); err != nil {
		fmt.Println("Div by zero:", err)
	}

	// --- Visibility demo ---
	fmt.Println("\n--- Visibility Rule ---")
	fmt.Println("calculator.Add  ← exported (uppercase A) — accessible")
	fmt.Println("calculator.Mul  ← exported (uppercase M) — accessible")
	// calculator.isZero(5)  ← compile error: unexported function
	fmt.Println("calculator.isZero ← unexported (lowercase i) — NOT accessible from here")
}
