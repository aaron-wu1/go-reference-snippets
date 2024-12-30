package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Walking(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	defer close(ch) // close channel after Walk() finishes
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	x := make(chan int)
	y := make(chan int)

	go Walking(t1, x)
	go Walking(t2, y)

	// check from channel
	for {
		v1, ok1 := <-x
		v2, ok2 := <-y
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		// if channel is closed, break
		if !ok1 {
			break
		}
		fmt.Println("v1: ", v1, "v2: ", v2)
	}
	return true
}

func main() {
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("Yes!")
	} else {
		fmt.Println("No!")
	}
}
