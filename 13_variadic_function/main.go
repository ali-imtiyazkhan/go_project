package main

import "fmt"

func sumAll(num ...int) int {

	total := 0

	for _, val := range num {
		total += val
	}

	return total
}

func main() {
	// sum := sumAll(1, 2, 3, 4, 5, 6)
	// fmt.Println(sum)

	// values := []int{10, 20, 30}

	// fmt.Println(sumAll(values...))

	// IIFE function

	res := func(a int, b int) int {
		return a * b
	}(10, 20)

	fmt.Println(res)
}
