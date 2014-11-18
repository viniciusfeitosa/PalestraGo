package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s \n", "World")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("App on :8999")
	log.Fatal(http.ListenAndServe(":8999", nil))
}
