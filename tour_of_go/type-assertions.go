package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string) // type assertion that i holds concrete type string and assigns it to s
	fmt.Println(s)

	s, ok := i.(string) // Similar to the above, but a test implementation
	fmt.Println(s, ok)  // This is a safe assertion

	f, ok := i.(float64) // Safe assertion on a wrong type, will work and ok should be false
	fmt.Println(f, ok)   // 0 false, so checks should happen on ok variable

	f = i.(float64) // panics because the assertion is unsafe
	fmt.Println(f)

}
