// https://go.dev/tour/flowcontrol/8
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// Validate CLI args
	if len(os.Args) < 3 {
		fmt.Println("Usage script.go x iters")
		os.Exit(1)
	}
	// Parse CLI args
	x, _ := strconv.ParseFloat(os.Args[1], 64)
	iters, _ := strconv.Atoi(os.Args[2])
	// Get sqrt accouting for errors
	res, err := Sqrt(x, iters)
	if err != nil {
		fmt.Println("Error when computing Sqrt:", err)
		os.Exit(1)
	}

	fmt.Println("Custom is:", *res, "math.Sqrt is:", math.Sqrt(x))
}

func Sqrt(x float64, iters int) (*float64, error) {
	// Check for negative input
	if x < 0 {
		return nil, ErrNegativeSqrt(x)
	}
	// Repeat how computers calcute square roots
	z := 1.0
	for i := 0; i < iters; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return &z, nil
}

// Custom error as per the exercises
type ErrNegativeSqrt float64

// Implemeting error interface by adding Error() func returning a string
func (e ErrNegativeSqrt) Error() string {
	// Because e is a float, we can Spint it
	// Not doing a conversion via float64() would result in an infinitelyr recursive call to Error()
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
	// %g is used for exponents and is default for floats
}
