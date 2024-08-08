package main

import (
	"log"
	"net/http"
)

func main() {
	loadEnv()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Printf("Starting server on %s", PORT)
	err := http.ListenAndServe(PORT, mux)
	log.Fatal(err)
}
