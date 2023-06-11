package main

import (
	"fmt"
	"math"
)

// Control our float precision
const FLOAT_PRECISION float64 = 3

// Define our custom structs
type square struct {
	sideLength float64
}
type triangle struct {
	base   float64
	height float64
}

// Define our interface
type shape interface {
	getArea() float64
}

func main() {
	// Some test shapes
	mySquare := square{sideLength: 10}
	myTriangle := triangle{
		base:   10,
		height: 10,
	}
	// Print our areas
	fmt.Println("Square area is:")
	printArea(mySquare)

	fmt.Println("Triangle area is:")
	printArea(myTriangle)

}

// Rounding helper
func roundToNearest(num, precision float64) float64 {
	scale := math.Pow(10, precision)
	return math.Round(num*scale) / scale
}

// Receiver funcs for our structs
// Starting with square
func (sq square) getArea() float64 {
	return roundToNearest(math.Pow(sq.sideLength, 2), FLOAT_PRECISION)
}

// Now a triangle
func (tr triangle) getArea() float64 {
	return roundToNearest(0.5*tr.base*tr.height, FLOAT_PRECISION)
}

// Interface function for printing area
func printArea(s shape) {
	fmt.Println(s.getArea())
}
