package main

import "fmt"

func main() {
	fmt.Println("here we go again")

	views := [5]int{1, 2, 3, 4, 5}

	sum := 0

	// Simple for loop
	for i := 0; i < len(views); i++ {
		sum += views[i]
	}

	total := 0

	// for Range Loop
	for i, v := range views {
		fmt.Println("day", i, "views", v)
		total = total + v
	}

	fmt.Println(total)

	fmt.Println("Sum:", sum)
}
