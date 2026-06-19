package main

import "fmt"

func main() {
	var marks [3]int

	marks[0] = 10
	marks[1] = 10
	marks[2] = 10

	fmt.Println(marks)

	// array literal

	res := [5]int{1, 2, 3, 4, 5}

	fmt.Println(res)
	fmt.Println(len(res))
}
