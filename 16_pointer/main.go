package main

import "fmt"

func main() {
	var a int = 10
	var p *int = &a

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Pointer p holds:", p)
	fmt.Println("Value via pointer:", *p)

	*p = 20
	fmt.Println("New value of a after pointer change:", a)

	// pointer with function
	b := 30
	zeroVal(b)
	fmt.Println("b after zeroVal:", b)

	zeroPtr(&b)
	fmt.Println("b after zeroPtr:", b)
}

func zeroVal(val int) {
	val = 0
}

func zeroPtr(ptr *int) {
	*ptr = 0
}
