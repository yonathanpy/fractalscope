package main

import (
	"fmt"
)

func main() {
	fmt.Println("[fractalscope] system boot")

	for i := 0; i < 15; i++ {
		fmt.Println("[cycle]", i)
	}

	fmt.Println("[fractalscope] system halt")
}
