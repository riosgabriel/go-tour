package main
import (
	"strconv"
	"fmt"
)

func main() {

	i, err := strconv.Atoi("abc")

	if err != nil {
		fmt.Printf("Não foi possível converter o número: %v\n", err)
		return
	}

	fmt.Println("Número convertido: ", i)
}
