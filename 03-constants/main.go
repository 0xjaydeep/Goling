// Demonstrates: const keyword, typed vs untyped constants,
// iota for enums (weekdays), and iota with bit-shift expressions (KB/MB/GB).
package main

import "fmt"

// --- Typed constant: type is fixed ---
const MaxRetries int = 3

// --- Untyped constant: adapts to the context it is used in ---
const Pi = 3.14159265358979

// --- iota: auto-incrementing integer, resets to 0 in each const block ---
type Weekday int

const (
	Sunday    Weekday = iota // 0
	Monday                   // 1
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

// String() makes Weekday satisfy fmt.Stringer — fmt.Println calls it automatically.
func (d Weekday) String() string {
	names := [...]string{
		"Sunday", "Monday", "Tuesday", "Wednesday",
		"Thursday", "Friday", "Saturday",
	}
	if d < Sunday || d > Saturday {
		return "Unknown"
	}
	return names[d]
}

// --- iota with bit-shift expressions ---
const (
	_  = iota             // skip 0 with blank identifier
	KB = 1 << (10 * iota) // 1 << 10 = 1 024
	MB                    // 1 << 20 = 1 048 576
	GB                    // 1 << 30 = 1 073 741 824
)

func main() {
	// --- Typed & untyped constants ---
	fmt.Println("--- Typed & Untyped Constants ---")
	fmt.Println("MaxRetries:", MaxRetries)
	fmt.Printf("Pi: %.10f\n", Pi)

	// Untyped constant adapts: works as float64 in arithmetic
	radius := 5.0
	area := Pi * radius * radius
	fmt.Printf("Circle area (r=5): %.4f\n", area)

	// --- Weekday enum with iota ---
	fmt.Println("\n--- Weekday Enum (iota) ---")
	today := Wednesday
	fmt.Println("Today:", today)          // prints "Wednesday" via String()
	fmt.Println("Numeric value:", int(today))

	fmt.Println("All days:")
	for d := Sunday; d <= Saturday; d++ {
		fmt.Printf("  %d → %s\n", d, d)
	}

	// --- Byte-size constants with iota + bit-shift ---
	fmt.Println("\n--- Byte Sizes (iota + bit-shift) ---")
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)

	fileSize := 2*GB + 512*MB
	fmt.Printf("2 GB + 512 MB = %d bytes\n", fileSize)
}
