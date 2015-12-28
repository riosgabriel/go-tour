package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Running on: ")
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Print("Nooooooooooooo.")
	case "linux":
		fmt.Print("Linux")
	default:
		fmt.Print("%s", os)
	}
}
