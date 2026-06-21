package main

import "fmt"

func main() {
	defer fmt.Println("This runs last (deferred)")

	fmt.Println("This runs first")

	defer fmt.Println("This runs second (deferred)")

	fmt.Println("This runs second")

	// defer with a function
	defer func() {
		fmt.Println("Anonymous deferred function runs at the end")
	}()
}
