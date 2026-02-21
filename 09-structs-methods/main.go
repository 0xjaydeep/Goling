// Demonstrates: struct definition, named and zero-value init, pointer to struct,
// embedded struct (composition), anonymous struct, value receiver method,
// pointer receiver method that modifies state, and String() (fmt.Stringer).
package main

import "fmt"

// --- Basic struct ---
type Point struct {
	X, Y float64
}

// --- Embedded struct (composition — Go has no class inheritance) ---
type Circle struct {
	Center Point
	Radius float64
}

// --- Struct with multiple fields ---
type Person struct {
	Name string
	Age  int
	City string
}

// --- Value receiver: operates on a copy, cannot modify original ---
func (p Person) Greet() string {
	return fmt.Sprintf("Hi, I'm %s from %s, age %d.", p.Name, p.City, p.Age)
}

// --- Pointer receiver: operates on original, can modify fields ---
func (p *Person) Birthday() {
	p.Age++ // modifies the actual Person, not a copy
}

// --- String() satisfies fmt.Stringer — fmt.Println calls it automatically ---
func (p Person) String() string {
	return fmt.Sprintf("Person{%s, %d, %s}", p.Name, p.Age, p.City)
}

func (c Circle) Area() float64 {
	return 3.14159265 * c.Radius * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle{center:(%.1f,%.1f) r:%.1f}", c.Center.X, c.Center.Y, c.Radius)
}

func main() {
	// --- Named initialisation (preferred — explicit field names) ---
	fmt.Println("--- Struct Init ---")
	p1 := Person{Name: "Alice", Age: 30, City: "NYC"}
	fmt.Println(p1) // calls String() automatically

	// --- Zero-value init (every field gets its zero value) ---
	var p2 Person
	fmt.Println("zero Person:", p2)

	// --- Field access ---
	p2.Name = "Bob"
	p2.Age = 25
	p2.City = "London"
	fmt.Println("updated p2:", p2)

	// --- Pointer to struct ---
	fmt.Println("\n--- Pointer to Struct ---")
	p3 := &Person{Name: "Carol", Age: 28, City: "Paris"}
	fmt.Println(p3)       // fmt knows how to dereference pointers
	fmt.Println(p3.Name)  // Go auto-dereferences: same as (*p3).Name

	// --- Methods ---
	fmt.Println("\n--- Methods ---")
	fmt.Println(p1.Greet())
	fmt.Printf("Before birthday: age=%d\n", p1.Age)
	p1.Birthday() // Go automatically takes &p1 because Birthday has pointer receiver
	fmt.Printf("After birthday:  age=%d\n", p1.Age)

	// --- Embedded struct (composition) ---
	fmt.Println("\n--- Embedded Struct (Composition) ---")
	c := Circle{Center: Point{3, 4}, Radius: 5}
	fmt.Println(c)
	fmt.Printf("Area: %.4f\n", c.Area())
	fmt.Printf("Center: (%.1f, %.1f)\n", c.Center.X, c.Center.Y)

	// --- Anonymous struct (one-off data, no type declaration needed) ---
	fmt.Println("\n--- Anonymous Struct ---")
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Server: %s:%d\n", config.Host, config.Port)
}
