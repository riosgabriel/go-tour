package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func main() {

	walk(tree.New(1))
}

func walk(t *tree.Tree) {
	if t == nil {
		return
	}

	walk(t.Left)
	fmt.Println(t.Value)
	walk(t.Right)
}
