package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	content := flag.String("content", "Hello, World!", "Content to display")
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, *content)
	})

	log.Printf("Starting server on port %s...", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
