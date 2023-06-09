package main

import "fmt"

// Declaring types
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreetingSb(sb)

}

func (englishBot) getGreeting() string {
	// We don't actually pass a variable because we don't use it
	// Very cusom logic for Eglish greeting
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	// We don't actually pass a variable because we don't use it[2]
	// Very cusom logic for Spanish greeting
	return "Ola!"
}

// Kinda duplicate function with different input type
// Not implementable this way
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
