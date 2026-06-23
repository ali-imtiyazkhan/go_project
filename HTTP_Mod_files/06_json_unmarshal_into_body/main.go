package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	url := "https://catfact.ninja/fact"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var fact CatFact
	err = json.Unmarshal(body, &fact)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Fact: %s\nLength: %d\n", fact.Fact, fact.Length)
}
