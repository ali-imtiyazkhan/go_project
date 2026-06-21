package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{Name: "Ali", Age: 25}
	fmt.Println("p1:", p1)

	p2 := Person{"John", 30}
	fmt.Println("p2:", p2)

	fmt.Println("p1.Name:", p1.Name)

	p1.Age = 26
	fmt.Println("p1 after age change:", p1)
}
