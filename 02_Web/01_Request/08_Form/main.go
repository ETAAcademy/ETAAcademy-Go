package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// r.Form will display both the values from the URL and the form if they have the same key
		fmt.Fprintln(w, r.PostForm)
	})
	server.ListenAndServe()
}
