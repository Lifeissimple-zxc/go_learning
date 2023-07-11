package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// Computes x to nth power & compares with lim
// If v < lim, return v else return lim
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v // v is limited to body of this if scope-wise!
		// But v can also be used in an else statement following the IF
	}
	return lim
}
