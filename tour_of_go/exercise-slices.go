package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// Create a slice of size dy
	outer := make([][]uint8, dy)
	// Make each element there a slice of dx 8bit unsigned ints
	for i := range outer {
		// Create inner slice on dx 8bit unsigned ints
		inner := make([]uint8, dx)
		outer[i] = inner
	}
	return outer

}

func main() {
	mySlice := Pic(2, 6)
	fmt.Println(mySlice)
	pic.Show(Pic(2, 6))
}
