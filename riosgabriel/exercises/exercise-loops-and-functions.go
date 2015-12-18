package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	previous := 0.0

	for i := 1; i <= 20; i++ {
		z = z - ((z * z - x) / 2*z)

		if previous == z {
			return z
		}

		previous = z
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
