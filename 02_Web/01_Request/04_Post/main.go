package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		// Print the headers
		fmt.Fprintln(w, r.Header)

		// Print the body
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body)) // The result is written to ResponseWriter, needs to be converted to a string to be readable
	})
	server.ListenAndServe()
}
