package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Base case when a tree is null
	if t == nil {
		return
	}
	ch <- t.Value // send value to the channel
	// Recursive calls to check left and right nodes
	Walk(t.Left, ch)
	Walk(t.Right, ch)
	// This needs to be closed somehow before we read from the channel?
	// Closing from within this function is not possible
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// Start by creating channels for async Walk calls
	ch1, ch2 := make(chan int), make(chan int)
	// Splitting walks across 2 goros for concurrency
	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()
	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()
	// With both trees traversed and channels closed, we can compare channels data
	seq1, seq2 := []int{}, []int{}
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 {
			seq1 = append(seq1, v1)
		}
		if ok2 {
			seq2 = append(seq2, v2)
		}
		// For equal sequences we should get same ok flags
		if ok1 != ok2 {
			return false
		}
		// vakues are same we don't check here
		// break the loop if both are not ok
		if !ok1 && !ok2 {
			break
		}
	}
	// // Different len means different sequences
	// if len(seq1) != len(seq2) {
	// 	return false
	// }
	// Getting here we can expect the sequences to have same len
	sort.Ints(seq1)
	sort.Ints(seq2)
	for i := 0; i < len(seq1); i++ {
		if seq1[i] != seq2[i] {
			return false
		}
	}
	return true
}

func main() {
	resSame := Same(tree.New(1), tree.New(1))
	fmt.Println("Comparing a pair of tree.New(1):", resSame)

	resDif := Same(tree.New(1), tree.New(2))
	fmt.Println("Comparing different trees:", resDif)

}
