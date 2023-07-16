package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Prepare result container
	res := make(map[string]int)
	// Split a string into words
	wrds := strings.Fields(s)
	for _, w := range wrds {
		res[w]++
	}
	return res
}

func main() {
	wc.Test(WordCount)
	fmt.Println(WordCount("car car car bike cheers"))
}
