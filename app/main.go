package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("success runned")
	http.HandleFunc("/hi", xxx())
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Printf("err %v", err)
	}
}

func xxx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("meow")
		word := os.Getenv("WHAT_SAY")
		if word != "" {
			respond(w, 200, word)
			return
		}

		respond(w, 404, "word not found")
	}
}

func respond(w http.ResponseWriter, code int, data string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
