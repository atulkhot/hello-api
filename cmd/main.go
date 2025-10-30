package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/hello", func(w http.ResponseWriter, r *http.Request) {
			enc := json.NewEncoder(w)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			resp := Resp{
				Language:   "English",
				Traslation: "Hello",
			}
			if err := enc.Encode(resp); err != nil {
				panic(fmt.Sprintf("Unable to encode response: %v", err))
			}
		})

	log.Printf("Listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language   string `json:"language"`
	Traslation string `json:"translation"`
}
