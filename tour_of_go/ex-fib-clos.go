package main

import "fmt"

func fibo() func() int {
	a, b := 0, 1 // Base scenario for fib sequence
	return func() int {
		a, b = b, a+b
		return b
	}

}

func main() {
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
