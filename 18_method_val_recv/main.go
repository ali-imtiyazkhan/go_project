package main

import "fmt"

type Counter struct {
	Value int
}

func (c Counter) Increment() {
	c.Value++
}

func main() {
	c := Counter{Value: 10}
	fmt.Println("Before:", c.Value)

	c.Increment()
	fmt.Println("After:", c.Value)
}
