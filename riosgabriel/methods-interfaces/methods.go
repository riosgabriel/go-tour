package main

import "fmt"

func main() {

	rec := &Rectangle{3, 5}

	fmt.Println("Area:", rec.area())
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) area()int {
	return r.width * r.height
}



