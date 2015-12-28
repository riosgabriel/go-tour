package main
import "fmt"

func main() {

	p := Person{27, "Gabriel"}

	fmt.Println(p.Age)
	fmt.Println(p.Name)
}

type Person struct {
	Age int
	Name string
}