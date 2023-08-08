package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	enemies := []string{"Vergil", "Nero", "Lady"}
	wg.Add(len(enemies))
	for _, e := range enemies {
		go attack(e, &wg)
	}
	wg.Wait()

	fmt.Println("Mission Accomplished")
}

func attack(enemy string, wg *sync.WaitGroup) {
	fmt.Printf("Dante attacked %v\n", enemy)
	wg.Done()
}
