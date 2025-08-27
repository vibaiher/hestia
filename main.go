package main

import (
	"fmt"
	"log"
	"net/http"

	"hestia/handlers"
)

func main() {
	http.Handle("/", handlers.Home())
	
	fmt.Printf("Starting Hestia server at http://localhost:8080\n")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}