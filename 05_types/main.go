package main

import (
	"fmt"
)

func main() {
	fmt.Println("here we go again")

	firstName := "Imtiyaz"
	lastName := "Khan"

	fullName := firstName + " " + lastName

	fmt.Println("Full name of the author is:", fullName)
	fmt.Printf("Full name type is: %T\n", fullName)
}