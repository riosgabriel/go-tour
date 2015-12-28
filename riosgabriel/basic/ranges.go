package main
import "fmt"


func main() {

	arr := []int {1, 2, 3, 4, 5}

	for index, value := range arr {
		fmt.Printf("Index [%d], value [%d]\n", index, value)
	}

	pow := initSlice(make([]int, 10))

	printValuesSlice(pow)
}

func initSlice(a []int)([]int) {
	for i := range a {
		a[i] = i << uint(i)
	}

	return a
}

func printValuesSlice(a []int) {
	for _, v := range a {
		fmt.Println(v)
	}
}