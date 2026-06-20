package main

import "fmt"

func main() {
	fmt.Println("here we go again")

	views := [5]int{1, 2, 3, 4, 5}

	sum := 0

	for i := 0; i < len(views); i++ {
		sum += views[i]
	}

	fmt.Println("Sum:", sum)
}