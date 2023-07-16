package main

import "fmt"

func main() {
	i, j := 42, 2701

	ptrToi := &i                                            // Get pointer to i
	fmt.Println("Printing i's value from ptrToi:", *ptrToi) //42
	*ptrToi = 322                                           // Update value via a pointer
	fmt.Println("I after pointer update is:", i)            // 322

	ptrToj := &j                // Pointer to j
	*ptrToj = *ptrToi / 322     // Update j's value via a pointer
	fmt.Println("New j is:", j) // 1
}
