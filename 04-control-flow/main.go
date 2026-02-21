// Demonstrates: if/else, if with init statement (err := fn(); err != nil),
// switch on value with multiple cases per arm, switch without condition
// (as a cleaner if/else chain), and fallthrough.
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func classify(n int) string {
	if n < 0 {
		return "negative"
	} else if n == 0 {
		return "zero"
	} else {
		return "positive"
	}
}

func dayType(day string) string {
	switch day {
	case "Saturday", "Sunday": // multiple values in one case
		return "weekend"
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "weekday"
	default:
		return "unknown"
	}
}

// grade uses switch without a condition — cleaner than a long if/else chain.
func grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func main() {
	// --- Basic if/else ---
	fmt.Println("--- if/else ---")
	fmt.Println(classify(-5))
	fmt.Println(classify(0))
	fmt.Println(classify(42))

	// --- if with init statement ---
	// The result and err variables are scoped to the if/else block only.
	fmt.Println("\n--- if with init statement ---")
	if result, err := divide(10, 3); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	if _, err := divide(5, 0); err != nil {
		fmt.Println("Caught error:", err)
	}

	// --- switch on value ---
	fmt.Println("\n--- switch on value ---")
	fmt.Println("Saturday:", dayType("Saturday"))
	fmt.Println("Monday:  ", dayType("Monday"))
	fmt.Println("Holiday: ", dayType("Holiday"))

	// --- switch without condition (grades) ---
	fmt.Println("\n--- switch as if/else chain (grades) ---")
	for _, s := range []int{95, 83, 72, 61, 45} {
		fmt.Printf("Score %d → Grade %s\n", s, grade(s))
	}

	// --- fallthrough: forces execution of the NEXT case body ---
	// Rarely used — Go does NOT fall through by default (unlike C).
	fmt.Println("\n--- fallthrough demo ---")
	n := 1
	switch n {
	case 1:
		fmt.Println("case 1 matched")
		fallthrough // runs case 2 body even though n != 2
	case 2:
		fmt.Println("case 2 (reached via fallthrough)")
	case 3:
		fmt.Println("case 3")
	}
}
