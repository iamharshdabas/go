package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello, World! The time is %s", time.Now().Format(time.RFC1123)); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		if _, err := fmt.Fprintf(w, "ParseForm() err: %v", err); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	}

	log.Printf("Received form data: %v", r.Form)
	if _, err := fmt.Fprintf(w, "POST request successful\n"); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
