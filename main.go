package main

import (
	"log"
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	log.Println("SERVER init 20000")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":20000", nil)
}
