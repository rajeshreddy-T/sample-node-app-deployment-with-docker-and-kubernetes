// sample go server

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s", r.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
