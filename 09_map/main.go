package main

import "fmt"

func main() {
	fmt.Println("hii there")

	// map[keytype]valuetype

	ages := map[string]int{
		"imtiyaz": 20,
		"jayesh":  21,
		"kaps":    21,
	}

	fmt.Println(ages["imtiyaz"])
}
