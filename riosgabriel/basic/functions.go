package main

import "fmt"

func main() {
	fmt.Println(mult(add(13, 22), 5))

	fmt.Println(double(1, 2, 3))
}

func mult(x, y int) (z, w int) {
	z * w
	return
}

func add(x, y int) int {
	return x + y
}

func double(x, y, z int)(int, int, int) {
	return x*x, y*y, z*z
}
