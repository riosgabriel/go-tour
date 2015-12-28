package main

import "fmt"

type Power 	struct {
	Value int
}

func main() {

	m := map[string]Power {
		"Darth Vader" : {60},
		"Luke Skywalker" : Power{45},
	}

	m["Gabriel"] = Power{100}

	fmt.Println(m)

	value, exists := m["Darth Vader"]

	fmt.Println("The value:", value, "Exists?", exists)
}