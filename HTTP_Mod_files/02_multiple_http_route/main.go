package main

import (
	"fmt"
	"net/http"
)

func checkHealth(w http.ResponseWriter, r *http.Request) {

	_, _ = w.Write([]byte("welcome to try to /hello?name=imtiyaz"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}

	_, _ = w.Write([]byte(name))
}

func main() {

	http.HandleFunc("/", checkHealth)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
