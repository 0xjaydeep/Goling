// Demonstrates: interface definition with multiple methods, two concrete types
// implementing the interface, function accepting interface (polymorphism),
// type assertion with ok check, type switch, and the empty interface (any).
package main

import (
	"fmt"
	"math"
)

// --- Interface: satisfied implicitly — no "implements" keyword needed ---
type Shape interface {
	Area() float64
	Perimeter() float64
}

// --- Concrete type 1: Rectangle ---
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.1f × %.1f)", r.Width, r.Height)
}

// --- Concrete type 2: Circle ---
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }
func (c Circle) String() string {
	return fmt.Sprintf("Circle(r=%.1f)", c.Radius)
}

// --- Function that accepts any Shape (polymorphism) ---
func printShapeInfo(s Shape) {
	fmt.Printf("  Area:      %.4f\n", s.Area())
	fmt.Printf("  Perimeter: %.4f\n", s.Perimeter())
}

// --- Type switch: branch on the concrete type ---
func describe(s Shape) string {
	switch v := s.(type) {
	case Rectangle:
		return fmt.Sprintf("a Rectangle %.1f wide and %.1f tall", v.Width, v.Height)
	case Circle:
		return fmt.Sprintf("a Circle with radius %.1f", v.Radius)
	default:
		return fmt.Sprintf("an unknown shape: %T", v)
	}
}

func main() {
	// --- Polymorphism via interface ---
	fmt.Println("--- Interface Polymorphism ---")
	shapes := []Shape{
		Rectangle{Width: 4, Height: 6},
		Circle{Radius: 3},
		Rectangle{Width: 10, Height: 2},
	}
	for _, s := range shapes {
		fmt.Printf("%v\n", s)
		printShapeInfo(s)
	}

	// --- Type assertion (extract the concrete value) ---
	fmt.Println("\n--- Type Assertion ---")
	var s Shape = Circle{Radius: 5}

	if c, ok := s.(Circle); ok {
		fmt.Printf("It's a Circle — radius %.1f, area %.4f\n", c.Radius, c.Area())
	}
	if _, ok := s.(Rectangle); !ok {
		fmt.Println("Not a Rectangle")
	}

	// --- Type switch ---
	fmt.Println("\n--- Type Switch ---")
	for _, shape := range shapes {
		fmt.Println("It's", describe(shape))
	}

	// --- Empty interface: can hold any value ---
	fmt.Println("\n--- Empty Interface (any) ---")
	values := []interface{}{
		42, "hello", true, 3.14,
		Circle{Radius: 1}, Rectangle{Width: 2, Height: 3},
	}
	for _, v := range values {
		fmt.Printf("  %-25v  type: %T\n", v, v)
	}
}
