package main

import (
	"log"
	"net/http"
	"shipping_go/handlers/rest"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)
	
	log.Printf("Listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language   string `json:"language"`
	Traslation string `json:"translation"`
}
