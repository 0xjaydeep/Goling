// Demonstrates: basic function, shorthand params, multiple return values,
// named return values with naked return, variadic functions, spread operator (...),
// function as variable, higher-order function (apply pattern), and closures.
package main

import "fmt"

// --- Basic function ---
func add(a int, b int) int {
	return a + b
}

// --- Shorthand params: same type listed once ---
func multiply(a, b int) int {
	return a * b
}

// --- Multiple return values ---
func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// --- Named return values with naked return ---
// Names act as pre-declared variables; a bare return sends them back.
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return // naked return — returns named vars as-is
	}
	result = a / b
	return
}

// --- Variadic function: accepts zero or more ints ---
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// --- Higher-order function: accepts a function as parameter ---
func apply(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = fn(v)
	}
	return result
}

// --- Closure factory: each call gets its own independent counter ---
func makeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func main() {
	// Basic & shorthand params
	fmt.Println("--- Basic & Shorthand Params ---")
	fmt.Println("add(3, 4):", add(3, 4))
	fmt.Println("multiply(6, 7):", multiply(6, 7))

	// Multiple return values
	fmt.Println("\n--- Multiple Return Values ---")
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	min, max := minMax(nums)
	fmt.Printf("nums: %v\nmin: %d, max: %d\n", nums, min, max)

	// Named returns / naked return
	fmt.Println("\n--- Named Returns ---")
	if r, err := divide(10, 3); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", r)
	}
	if _, err := divide(5, 0); err != nil {
		fmt.Println("Caught:", err)
	}

	// Variadic
	fmt.Println("\n--- Variadic ---")
	fmt.Println("sum(1,2,3):", sum(1, 2, 3))
	fmt.Println("sum(1..5): ", sum(1, 2, 3, 4, 5))
	// Spread operator: unpack a slice into a variadic call
	values := []int{10, 20, 30}
	fmt.Println("sum(slice...):", sum(values...))

	// Function as variable
	fmt.Println("\n--- Function as Variable ---")
	double := func(n int) int { return n * 2 }
	fmt.Println("double(7):", double(7))
	fmt.Printf("type: %T\n", double)

	// Higher-order function
	fmt.Println("\n--- Higher-Order Function (apply) ---")
	data := []int{1, 2, 3, 4, 5}
	doubled := apply(data, func(n int) int { return n * 2 })
	squared := apply(data, func(n int) int { return n * n })
	fmt.Println("original:", data)
	fmt.Println("doubled: ", doubled)
	fmt.Println("squared: ", squared)

	// Closure
	fmt.Println("\n--- Closure (makeCounter) ---")
	counter := makeCounter(0)
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	other := makeCounter(10)
	fmt.Println(other())  // 11 — independent state
	fmt.Println(counter()) // 4  — counter keeps its own state
}
