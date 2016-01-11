package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	InnerWalk(t, ch)
	close(ch)
}

func InnerWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		InnerWalk(t.Left, ch)
		ch <- t.Value
		InnerWalk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v := range ch1 {
		if (v != <-ch2) {
			return false
		}
	}

	return true
}

func main() {

	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(3)))
}
