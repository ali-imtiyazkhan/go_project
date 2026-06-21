package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok":       true,
		"message":  "JSon encode successful",
		"dateTime": time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	var data map[string]any
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"ok":      true,
		"message": "JSON decode successful",
		"data":    data,
	})
}

func main() {

	http.HandleFunc("/ok", rootHandler)
	http.HandleFunc("/decode", decodeHandler)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
