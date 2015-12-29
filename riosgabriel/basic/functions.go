package main

import "fmt"

func main() {
	fmt.Println(mult(add(13, 22), 5))

	fmt.Println(double(1, 2, 3))

	fmt.Println(compute(1, 2, add))
}

func compute(x, y int, fn func(x int, y int) int) int {
	return fn(x, y)
}

func mult(x, y int)int {

	return x * y
}

func add(x, y int) int {
	return x + y
}

func double(x, y, z int)(int, int, int) {
	return x*x, y*y, z*z
}
