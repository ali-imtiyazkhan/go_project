package main

import "fmt"

func add(a int, b int) int {
	var c = a + b
	return c
}

func main() {
	a, b := 10, 20

	total := add(a, b)

	fmt.Println("sum of the a and b is : ", total)
}
