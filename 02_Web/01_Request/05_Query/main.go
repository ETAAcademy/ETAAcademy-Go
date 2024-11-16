package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query()

		// Method 1: Retrieve values through a map, returns a slice of strings
		id := query["id"]
		log.Println(id)

		// Method 2: Use the Get method, returns the first value for the given key
		name := query.Get("name")
		log.Println(name)

	})
	http.ListenAndServe("localhost:8080", nil)
}

// In the browser, input: http://localhost:8080/home?id=001&name=barbie&id=002&name=go
