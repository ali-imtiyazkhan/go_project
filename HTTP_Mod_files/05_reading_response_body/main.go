package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func okRoute(w http.ResponseWriter, r *http.Request) {

	res := map[string]any{
		"message": "all is right",
		"ok":      true,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}

func main() {

	url := "https://jsonplaceholder.typicode.com/todos"

	resp, err1 := http.Get(url)

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	http.HandleFunc("/ok", okRoute)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)

}
