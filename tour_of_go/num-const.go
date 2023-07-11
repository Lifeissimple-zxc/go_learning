package main

import "fmt"

const (
	// Create a huge number by shifting 1 bit left 100 places
	// We get a binary that is 1 followed by 100 zeroes
	Big = 1 << 100
	// Shift it right again 99 places so we end up with 1<<1, or 2
	Small = Big >> 99
)

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big)) // this overflows!
}

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }
