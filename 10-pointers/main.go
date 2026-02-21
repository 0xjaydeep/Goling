// Demonstrates: & (address-of), * (dereference), modifying through a pointer,
// pass-by-value vs pass-by-pointer, new(), pointer to struct, and nil pointer.
package main

import "fmt"

// --- Pass by value: receives a COPY, original is untouched ---
func doubleByValue(n int) int {
	n *= 2
	return n
}

// --- Pass by pointer: receives the address, modifies original ---
func doubleByPointer(n *int) {
	*n *= 2
}

// --- Increment via pointer ---
func increment(n *int) {
	*n++
}

func main() {
	// --- Basic pointer operators ---
	fmt.Println("--- & (address-of) and * (dereference) ---")
	x := 42
	p := &x               // p is *int — holds the memory address of x
	fmt.Println("x  =", x)
	fmt.Println("p  =", p)  // prints something like 0xc0000b4010
	fmt.Println("*p =", *p) // dereferences: the value stored at that address

	*p = 100 // change x through the pointer
	fmt.Println("x after *p = 100:", x)

	// --- Pass-by-value vs pass-by-pointer ---
	fmt.Println("\n--- Pass-by-Value vs Pass-by-Pointer ---")
	a := 7
	result := doubleByValue(a)
	fmt.Printf("doubleByValue:   a=%d (unchanged),  result=%d\n", a, result)

	b := 7
	doubleByPointer(&b) // pass the address of b
	fmt.Printf("doubleByPointer: b=%d (modified in place)\n", b)

	// --- Accumulating via pointer ---
	fmt.Println("\n--- Increment via Pointer ---")
	counter := 0
	for i := 0; i < 5; i++ {
		increment(&counter)
	}
	fmt.Println("counter after 5 increments:", counter)

	// --- new(): allocates zeroed memory, returns a pointer ---
	fmt.Println("\n--- new() ---")
	np := new(int) // *int pointing to a zero int
	fmt.Println("*new(int):", *np)
	*np = 55
	fmt.Println("after *np = 55:", *np)

	// --- Pointer to struct: Go auto-dereferences field access ---
	fmt.Println("\n--- Pointer to Struct ---")
	type Point struct{ X, Y int }
	pt := &Point{X: 1, Y: 2}
	pt.X = 10 // same as (*pt).X = 10 — Go adds the dereference automatically
	fmt.Println("Point:", *pt)

	// --- Nil pointer ---
	fmt.Println("\n--- Nil Pointer (do NOT dereference!) ---")
	var nilPtr *int
	fmt.Println("nil pointer:", nilPtr)
	fmt.Println("is nil:", nilPtr == nil)
	// The line below would cause a runtime panic — never dereference nil:
	// fmt.Println(*nilPtr)  // PANIC: runtime error: invalid memory address or nil pointer dereference
	fmt.Println("(dereferencing a nil pointer causes a panic — always check nil first)")
}
