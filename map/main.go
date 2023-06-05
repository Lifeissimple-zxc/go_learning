package main

import "fmt"

func main() {
	// Declaring a map of {name: hexcode} format
	// Option 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"black": "#000000",
	}

	// Option 2: more verbose
	// var colors map[string]string // this initialises a map with zero values
	// We use htis option when we don't know (or don't yet have)
	// The data to append to that map

	// Option 3: using make
	// colors := make(map[string]string) // this initialises a map with its zero values (so an empty map)

	// Adding data to a map (very familiar)
	colors["white"] = "#ffffff"

	// Delete data from a map by key
	// delete(colors, "white")
	// fmt.Println(colors)

	// Iterative prints
	printMap(colors)

}

// Function to iterate over key, value of a map
func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Hex code for", key, "is", value)
	}
}
