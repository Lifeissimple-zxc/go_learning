package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10
	iptr := unsafe.Pointer(&i) // pointer that can bypass GO's strong typing
	// Cast to an integer pointer and print out!
	fmt.Println((*int)(iptr))
	fmt.Println(*(*int)(iptr))     // Dereference our int
	fmt.Println(*(*float64)(iptr)) // Dereference our int to a float64
	// fmt.Println(*(*string)(iptr))  // Dereference our int to a string

	// Pointer arithmetic
	arr := []int{1, 123, 3, 4, 5}
	arrPtr := unsafe.Pointer(&arr[0])
	for x := 0; x < len(arr); x++ {
		nxt := (*int)(unsafe.Pointer(uintptr(arrPtr) + uintptr(x)*unsafe.Sizeof(arr[0])))
		fmt.Println(nxt, *nxt)
	}

}
