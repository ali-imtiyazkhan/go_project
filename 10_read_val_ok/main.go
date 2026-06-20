package main

import "fmt"

func main() {
	points := map[string]int{
		"a": 10,
		"b": 0, //valid value
	}

	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"]) //output c : 0

	valB, okB := points["b"]

	fmt.Println(valB, okB)

	valC, okC := points["c"]

	fmt.Println(valC, okC)

	prices := map[string]int{
		"xyz": 100,
		"abc": 200,
		"def": 300,
	}

	total := 0

	for _, price := range prices {
		total += price
	}

	fmt.Println(total)
}
