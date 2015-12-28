package main

import "fmt"

func main() {

	var a [2]string

	a[0] = "Gabriel"
	a[1] = "Rios"

	fmt.Println(a)

	b := []int {1, 2, 3, 4}

	fmt.Println(b)

	x := make([]int, 0)
	x = append(x, 1, 2, 3)

	fmt.Println(x)

	x = append(x, 11, 12)

	fmt.Println(x)

	fmt.Println(x[2:4])

}
