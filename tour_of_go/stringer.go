package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zetnik Daun", 32}
	fmt.Println(a, z) // Arthur Dent (42 years) Zetnik Daun (32 years) as per String() logic

}

// Implementing this enables fmt.Prinln to print data on the struct
// Reason: it satisfies Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
