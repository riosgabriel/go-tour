package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt2(x float64) (float64, error) {
	z := 1.0
	previous := 0.0

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	for i := 1; i <= 20; i++ {
		z = z - ((z * z - x) / 2*z)

		if previous == z {
			return z, nil
		}

		previous = z
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}
