package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// This is the base code
	// fmt.Println(strings.Join(os.Args[1:], " "))
	// 1.1 Modify the code to print name of the programm invoking it -- start with 0 arg
	fmt.Println("String join Implementation:")
	tString := time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	// Timing our execution for exercise 1.3
	printTimeDelta("String Join Implementation Took:", tString)
	fmt.Println("#####################")
	// 1.2 Modify the code to print index and value, line by line
	fmt.Println("Indexed loop implementation")
	tIndexLoop := time.Now()
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i], "at index", i)
	}
	printTimeDelta("Indexed loop took: ", tIndexLoop)
	fmt.Println("#####################")
	// 1.3 Time different implementations
	fmt.Println("Non string join implementation:")
	tInefficient1 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		// Change our separator after the first argument
		sep = " "
	}
	fmt.Println(s)
	printTimeDelta("Non string join implementation took: ", tInefficient1)
	fmt.Println("#####################")

}

// Computes time elapsed from start and prints it with msg
func printTimeDelta(msg string, start time.Time) {
	fmt.Println(msg, time.Now().Sub(start))
}

/*
	strings.Join takes microseconds whilst other implementations take miliseconds / seconds
*/
