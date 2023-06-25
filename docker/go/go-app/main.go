package main

import (
	"fmt"
	"log"
	"net/http"
)

func echoHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Go!!</h1>")
}

func main() {
	// controller
	http.HandleFunc("/", echoHello)
	// port
	log.Fatal(http.ListenAndServe(":8000", nil))
}
