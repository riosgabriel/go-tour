package main

import (
	"fmt"
)

func main() {

	defer stackDefers()

	defer fmt.Println("Go! ")
	defer fmt.Print("and ")

	fmt.Print("Hello ")
	fmt.Print("World ")


}

func stackDefers() {
	fmt.Println("counting")

	defer fmt.Println("done")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
