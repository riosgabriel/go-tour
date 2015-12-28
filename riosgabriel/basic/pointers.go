package  main
import "fmt"

func main() {

	i := 30

	fmt.Printf("Value of i: %d \n", i)

	p := &i

	fmt.Printf("Value of &i: %d \n", *p)

	*p = 22

	fmt.Printf("New value of &i: %d \n", *p)

	fmt.Printf("New value of i: %d", i)
}


