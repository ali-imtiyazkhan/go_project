package main

import "fmt"

func main() {
	// common collector type
	// dynmaic collector can grow
	// []type{...}

	// make([]T,len,cap)

	result := []int{1, 2, 3, 4, 5}
	result = append(result, 12)

	score := make([]int, 0, 5)
	score = append(score, 0, 2)

	fmt.Println(score)

	fmt.Println(result, result[0], len(result))

}
