// Demonstrates: package declaration, imports, Println vs Printf,
// format verbs (%s %d %f %v %T), and zero values for basic types.
package main

import "fmt"

func main() {
	// --- Basic output ---
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to Go!")

	// --- Printf with format verbs ---
	name := "Gopher"
	age := 10
	pi := 3.14159
	active := true

	fmt.Printf("Name:    %s\n", name)    // %s  → string
	fmt.Printf("Age:     %d\n", age)     // %d  → integer (decimal)
	fmt.Printf("Pi:      %f\n", pi)      // %f  → floating-point
	fmt.Printf("Pi:      %.2f\n", pi)    // %.2f → 2 decimal places
	fmt.Printf("Active:  %v\n", active)  // %v  → default format (any type)
	fmt.Printf("Type of age: %T\n", age) // %T  → type name

	// --- Sprintf: format into a string instead of printing ---
	msg := fmt.Sprintf("Hi, I'm %s and I'm %d years old.", name, age)
	fmt.Println(msg)

	// --- Zero values (what Go assigns when you declare but don't initialise) ---
	var i int
	var f float64
	var s string
	var b bool

	fmt.Println("\n--- Zero Values ---")
	fmt.Printf("int     zero value: %v\n", i)
	fmt.Printf("float64 zero value: %v\n", f)
	fmt.Printf("string  zero value: %q\n", s) // %q shows quotes → easy to spot empty string
	fmt.Printf("bool    zero value: %v\n", b)
}
