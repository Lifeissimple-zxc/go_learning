package main

// Ex 1.4: modify the code to print the names of all files in which each duplicated line occurs.
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a container for counting strings with key as string and int as value
	counts := make(map[string]map[string]int)
	// Get list of files to check for duplicate strings
	files := os.Args[1:]
	// Validate that we actually have files to work with
	// If we don't --> work with stdin
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// Process files
		for _, arg := range files {
			// Open file from the arguments
			f, err := os.Open(arg)
			// Check for open error
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// Count lines in file
			countLines(f, counts)
			f.Close()
		}
	}
	// Print contents of our map
	fmt.Printf("Duplicate strings:\n")
	printDuplicateStrings(counts)
}

/*
Helper that performs counting of strings.
*/
func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// Because of the nested structure of our map, we need to handle nil scenario
		if counts[input.Text()] == nil {
			// Func local container that counts occurences in a given file
			fileOccurences := make(map[string]int)
			// Count our occurence
			fileOccurences[f.Name()]++
			// Save in global counts
			counts[input.Text()] = fileOccurences
		} else {
			counts[input.Text()][f.Name()]++
		}

	}
}

/*
Helper that counts and prints duplicates from our counts variable
*/
func printDuplicateStrings(counts map[string]map[string]int) {
	for line, innerMap := range counts {
		occurences := 0
		for _, n := range innerMap {
			// Count all occurences accross files to know if it is a duplicate or not
			occurences += n
		}
		if occurences > 1 {
			// pring word: occurences
			fmt.Printf("\n%s:\t%d.\n", line, occurences)
			// print file details
			fmt.Println("File Occurences Breakdown:")
			for file, n := range innerMap {
				fmt.Printf("%s:\t%d\n", file, n)
			}
		}
	}
}
