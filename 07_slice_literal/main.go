package main

import "fmt"

func main() {
	// common collector type
	// dynmaic collector can grow
	// []type{...}

	result := []int{1, 2, 3, 4, 5}
	result = append(result, 12)

	fmt.Println(result, result[0], len(result))
}
