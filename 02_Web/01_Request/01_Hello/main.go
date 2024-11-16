package main

import "net/http"

func main() {
	// Call the adapter handler function with two parameters: 
	// an HTTP address and a handler function.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello gooooo"))
	})

	// Set up a web server with two parameters: 
	// a listening address (port) and a handler. 
	// By default, the handler is nil, which uses the default multiplexer (mux).
	http.ListenAndServe("localhost:8080", nil)
}
