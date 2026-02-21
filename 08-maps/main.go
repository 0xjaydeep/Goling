// Demonstrates: map literal, make, set/get, zero value for missing key,
// two-value form (val, ok), delete, range over map, len, and nested maps.
package main

import "fmt"

func main() {
	// --- Map literal ---
	fmt.Println("--- Map Literal ---")
	scores := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
	}
	fmt.Println("scores:", scores)
	fmt.Println("len:", len(scores))

	// --- make: preferred when building maps dynamically ---
	fmt.Println("\n--- make(map) ---")
	inventory := make(map[string]int)
	inventory["apples"] = 10
	inventory["bananas"] = 5
	inventory["oranges"] = 8
	fmt.Println("inventory:", inventory)

	// --- Get: missing key returns the zero value (0 for int) ---
	fmt.Println("\n--- Zero Value for Missing Key ---")
	fmt.Println("Alice:", scores["Alice"]) // 95
	fmt.Println("Dave: ", scores["Dave"])  // 0 — does NOT error, just returns zero!

	// --- Two-value form: distinguish "not found" from "stored zero" ---
	fmt.Println("\n--- Two-Value Form (val, ok) ---")
	if val, ok := scores["Alice"]; ok {
		fmt.Println("Found Alice:", val)
	}
	if val, ok := scores["Dave"]; !ok {
		fmt.Printf("Dave not found (zero val = %d)\n", val)
	}

	// --- delete ---
	fmt.Println("\n--- delete ---")
	fmt.Println("Before:", scores)
	delete(scores, "Bob")
	delete(scores, "Nobody") // deleting a non-existent key is safe (no-op)
	fmt.Println("After deleting Bob:", scores)

	// --- range over map (iteration order is randomised every run) ---
	fmt.Println("\n--- range over map ---")
	for name, score := range scores {
		fmt.Printf("  %-6s → %d\n", name, score)
	}

	// --- Nested map ---
	fmt.Println("\n--- Nested Map ---")
	school := map[string]map[string]int{
		"Alice": {"Math": 95, "English": 88},
		"Bob":   {"Math": 72, "English": 81},
	}
	// Direct nested access
	fmt.Println("Alice's Math:", school["Alice"]["Math"])

	// Range outer, then inner
	for student, subjects := range school {
		fmt.Println(student + ":")
		for subject, grade := range subjects {
			fmt.Printf("  %-10s %d\n", subject, grade)
		}
	}
}
