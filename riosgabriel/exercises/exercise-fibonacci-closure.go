package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := 0

	return func() int {
		i := 1
		j := 0

		for x:= 0; x < n; x++ {
			t := i + j
			i = j
			j = t
		}

		n++
		return j
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
