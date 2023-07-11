// https://go.dev/tour/flowcontrol/8
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func Sqrt(x float64, iters int) float64 {
	// Repeat how computers calcute square roots
	z := 1.0
	for i := 0; i < iters; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	// Validate CLI args
	if len(os.Args) < 3 {
		fmt.Println("Usage script.go x iters")
		os.Exit(1)
	}
	// Parse CLI args
	x, _ := strconv.ParseFloat(os.Args[1], 64)
	iters, _ := strconv.Atoi(os.Args[2])
	fmt.Println("Custom is:", Sqrt(x, iters), "math.Sqrt is:", math.Sqrt(x))
}
