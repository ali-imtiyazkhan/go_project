package main

import "fmt"

func divide(a int, b int) (imtiyaz int, john int) {
	imtiyaz = a / b
	john = a + b
	return
}

func main() {
	divide1, add := divide(10, 20)

	fmt.Println(divide1, add)
}
