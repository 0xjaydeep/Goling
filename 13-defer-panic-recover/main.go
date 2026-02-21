// Demonstrates: defer execution order (LIFO), defer for cleanup (simulated file),
// panic + recover in a deferred function (safeDiv), and panic propagation
// through the call stack caught by a wrapper.
package main

import "fmt"

// --- defer for cleanup: the close always runs, even if the function errors ---
func readFile(name string) {
	fmt.Printf("  Opening  %s\n", name)
	defer fmt.Printf("  Closing  %s  (deferred cleanup)\n", name)
	fmt.Printf("  Reading  %s\n", name)
	// Imagine reading content here; defer ensures Close is called on exit.
}

// --- recover in a deferred func turns a panic into a returned error ---
func safeDiv(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered panic: %v", r)
		}
	}()
	result = a / b // integer division by zero is a runtime panic
	return
}

// --- panic propagates up the call stack ---
func riskyOperation() {
	fmt.Println("  riskyOperation: about to panic")
	panic("something went very wrong")
}

func safeWrapper() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("safeWrapper caught: %v", r)
		}
	}()
	riskyOperation()
	return nil // never reached when panic occurs
}

func main() {
	// --- LIFO (last-in, first-out) order ---
	fmt.Println("--- defer LIFO Order ---")
	defer fmt.Println("  defer 1 — runs LAST  (registered first)")
	defer fmt.Println("  defer 2 — runs second")
	defer fmt.Println("  defer 3 — runs FIRST  (registered last)")
	fmt.Println("  (main body — defers run after this block finishes)")
	fmt.Println()

	// --- defer for cleanup ---
	fmt.Println("--- defer for Cleanup ---")
	readFile("data.txt")
	fmt.Println()

	// --- panic + recover (safeDiv) ---
	fmt.Println("--- panic + recover (safeDiv) ---")
	if r, err := safeDiv(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", r)
	}
	if r, err := safeDiv(10, 0); err != nil {
		fmt.Println("Caught panic:", err)
	} else {
		fmt.Println("result:", r)
	}
	fmt.Println()

	// --- panic propagation caught by wrapper ---
	fmt.Println("--- panic propagation + recover in wrapper ---")
	if err := safeWrapper(); err != nil {
		fmt.Println("safeWrapper returned:", err)
	}
	fmt.Println("Program continues normally after recovered panic")
	fmt.Println()

	// The three defers from the top of main() fire here in LIFO order:
}
