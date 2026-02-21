// Demonstrates: three declaration styles, zero values for all basic types,
// multiple assignment (and swap), explicit type conversion, and fmt.Sprintf
// for converting an int to a string.
package main

import "fmt"

func main() {
	// --- Style 1: var with explicit type ---
	var x int = 42
	var greeting string = "Hello"
	fmt.Println("Style 1:", x, greeting)

	// --- Style 2: var with inferred type (compiler figures it out) ---
	var score = 98.6 // inferred float64
	var found = true // inferred bool
	fmt.Println("Style 2:", score, found)

	// --- Style 3: short declaration := (only inside functions) ---
	count := 7
	message := "Go is great"
	fmt.Println("Style 3:", count, message)

	// --- Zero values (declared but not initialised) ---
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("\n--- Zero Values ---")
	fmt.Printf("int:     %v\n", a)
	fmt.Printf("float64: %v\n", b)
	fmt.Printf("string:  %q\n", c) // %q adds quotes — helps spot the empty string
	fmt.Printf("bool:    %v\n", d)

	// --- Multiple assignment ---
	p, q := 10, 20
	fmt.Println("\n--- Multiple Assignment ---")
	fmt.Println("p:", p, "q:", q)
	p, q = q, p // idiomatic Go swap — no temp variable needed
	fmt.Println("After swap — p:", p, "q:", q)

	// --- Explicit type conversion (Go has NO implicit casting) ---
	var intVal int = 100
	var floatVal float64 = float64(intVal) // must be explicit
	var backToInt int = int(floatVal)      // truncates, does not round
	fmt.Println("\n--- Type Conversion ---")
	fmt.Println("int → float64:", floatVal)
	fmt.Println("float64 → int:", backToInt)

	// float with fractional part — int truncates it
	var pi float64 = 3.99
	fmt.Printf("float64(%.2f) → int(%d)  (truncates, not rounds)\n", pi, int(pi))

	// --- int → string: use fmt.Sprintf, NOT string(intVal) ---
	// string(100) gives 'd' (the rune for code point 100) — almost never what you want
	numStr := fmt.Sprintf("%d", intVal)
	fmt.Println("\nint → string via Sprintf:", numStr, "| type:", fmt.Sprintf("%T", numStr))
	fmt.Printf("string(100) gives: %q  ← rune, not the number!\n", string(100))
}
