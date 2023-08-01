// Based on a quick yt video: https://www.youtube.com/watch?v=fTKsaPiThwM
package main

import "fmt"

// Non-generic summer for an int array
func sumOfInt(intArr []int) int { // this only works for ints, but not for floats
	inc := 0
	for _, val := range intArr {
		inc += val
	}
	return inc
}

// Generic implementation
type Number interface {
	int64 | float64 // take either int64 or float64
}

// Generic: T is some type that we refer to after func name within [] braces
func sumOf[T Number](someArr []T) T {
	var inc T // initialise a variable of type T to compute our sum
	for _, val := range someArr {
		inc += val
	}
	return inc

}

func main() {
	fmt.Println("Non-generic:")
	fmt.Printf("%d\n", sumOfInt([]int{0, 1, 2, 3}))

	fmt.Println("Generic with an int:")
	fmt.Printf("%d\n", sumOf([]int64{0, 1, 2, 3}))

	fmt.Println("Generic with a float:")
	fmt.Printf("%0.2f\n", sumOf([]float64{0.0, 1.1, 2.2, 3.3}))
}
