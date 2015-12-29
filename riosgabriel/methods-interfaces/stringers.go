package main

import "fmt"

type Car struct {
	Model string
	Year string
}

func (car Car) String() string {
	return fmt.Sprintf("Model: %s, Year: %s", car.Model, car.Year)
}

func main() {
	c := Car{"Ford", "1985"}

	fmt.Println(c)
}