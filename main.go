package main

import (
	"gymshark/handlers"
	"log"
	"net/http"
)

func main() {
	http.Handle("/amount", http.HandlerFunc(handlers.Amount))
	http.Handle("/packs", http.HandlerFunc(handlers.Packs))

	log.Println("listening on port 8081")
	http.ListenAndServe(":8081", nil)
}
