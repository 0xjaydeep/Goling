// Demonstrates: traditional for, while-style for, infinite for with break,
// range over slice (index + value), range with _ to skip index,
// range over map, range over string (runes), and continue.
package main

import "fmt"

func main() {
	// --- Traditional for (C-style: init; condition; post) ---
	fmt.Println("--- Traditional for ---")
	for i := 0; i < 5; i++ {
		fmt.Printf("  i = %d\n", i)
	}

	// --- While-style for (condition only) ---
	fmt.Println("\n--- While-style for ---")
	n := 1
	for n < 32 {
		fmt.Printf("  n = %d\n", n)
		n *= 2
	}

	// --- Infinite for with break ---
	fmt.Println("\n--- Infinite for with break ---")
	sum := 0
	counter := 1
	for {
		sum += counter
		counter++
		if counter > 5 {
			break
		}
	}
	fmt.Println("  Sum 1..5:", sum)

	// --- range over slice: gives index and value ---
	fmt.Println("\n--- range over slice (index + value) ---")
	fruits := []string{"apple", "banana", "cherry"}
	for i, v := range fruits {
		fmt.Printf("  [%d] = %s\n", i, v)
	}

	// --- range with _ to discard the index ---
	fmt.Println("\n--- range, skip index with _ ---")
	for _, fruit := range fruits {
		fmt.Println(" ", fruit)
	}

	// --- continue: skip to next iteration ---
	fmt.Println("\n--- continue (skip even numbers) ---")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("  odd: %d\n", i)
	}

	// --- range over map (order is not guaranteed!) ---
	fmt.Println("\n--- range over map ---")
	scores := map[string]int{"Alice": 95, "Bob": 87, "Carol": 92}
	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}

	// --- range over string iterates Unicode code points (runes), not bytes ---
	fmt.Println("\n--- range over string (runes) ---")
	for i, r := range "Go! 🐹" {
		fmt.Printf("  byte index %d: %c  (U+%04X)\n", i, r, r)
	}
}
