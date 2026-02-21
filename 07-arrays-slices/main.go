// Demonstrates: arrays (value type / copy independence), slice literals,
// make with len/cap, append, slice expressions (a[1:3]), the shared
// backing-array gotcha, copy() for true independence, and nil slices.
package main

import "fmt"

func main() {
	// --- Arrays ---
	fmt.Println("--- Arrays (value type) ---")
	var arr [3]int              // zero-value: [0 0 0]
	arr2 := [3]int{10, 20, 30} // array literal
	arr3 := [...]int{1, 2, 3, 4} // [...] lets compiler count elements
	fmt.Println("arr:", arr)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr3:", arr3, "  len:", len(arr3))

	// Arrays are VALUE types — assignment copies the whole array
	copyArr := arr2
	copyArr[0] = 999
	fmt.Println("arr2 after modifying copyArr:", arr2) // unchanged!
	fmt.Println("copyArr:", copyArr)

	// --- Slice literals ---
	fmt.Println("\n--- Slice Literals ---")
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("s: %v  len=%d  cap=%d\n", s, len(s), cap(s))

	// make(type, length, capacity)
	s2 := make([]int, 3, 5)
	fmt.Printf("make([]int, 3, 5): %v  len=%d  cap=%d\n", s2, len(s2), cap(s2))

	// append: adds elements; may allocate a new backing array when cap exceeded
	s2 = append(s2, 100, 200)
	fmt.Printf("after 2 appends:   %v  len=%d  cap=%d\n", s2, len(s2), cap(s2))
	s2 = append(s2, 300) // exceeds original cap=5 → new allocation, cap grows
	fmt.Printf("after 3rd append:  %v  len=%d  cap=%d\n", s2, len(s2), cap(s2))

	// --- Slice expressions ---
	fmt.Println("\n--- Slice Expressions ---")
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("a:       ", a)
	fmt.Println("a[1:4]:  ", a[1:4]) // elements at index 1, 2, 3
	fmt.Println("a[:3]:   ", a[:3])  // first three
	fmt.Println("a[3:]:   ", a[3:])  // from index 3 onward

	// --- Shared backing array GOTCHA ---
	fmt.Println("\n--- Shared Backing Array Gotcha ---")
	original := []int{1, 2, 3, 4, 5}
	slice1 := original[1:4] // {2, 3, 4} — shares the same memory!
	slice1[0] = 999
	fmt.Println("original after modifying slice1:", original) // 2 → 999 !!
	fmt.Println("slice1:", slice1)

	// --- copy() creates independence ---
	fmt.Println("\n--- copy() for Independence ---")
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	n := copy(dst, src) // returns number of elements copied
	dst[0] = 777
	fmt.Println("src:", src) // unchanged
	fmt.Println("dst:", dst)
	fmt.Printf("copied %d elements\n", n)

	// --- Nil slice ---
	fmt.Println("\n--- Nil Slice ---")
	var nilSlice []int
	fmt.Println("nil slice:", nilSlice)
	fmt.Println("nil?", nilSlice == nil)
	fmt.Printf("len=%d  cap=%d\n", len(nilSlice), cap(nilSlice))
	nilSlice = append(nilSlice, 1, 2, 3) // append works perfectly on nil
	fmt.Println("after append:", nilSlice)
}
