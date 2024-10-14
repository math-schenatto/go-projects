package main

import (
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":5000", mux))
}
