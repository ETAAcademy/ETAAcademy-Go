package main

import "net/http"

func main() {
	mh := myHandler{} // This should be a pointer
	about := aboutHandler{}
	hello := helloHandler{}

	// http.ListenAndServe("localhost:8080", nil)

	// Equivalent to the above line
	server := http.Server{
		Addr:    ("localhost:8080"),
		Handler: nil,
		// Handler: &mh,
	}
	// Different paths are handled by different handlers
	http.Handle("/hello", &hello)
	http.Handle("/about", &about)
	http.Handle("/home", &mh)

	server.ListenAndServe()
}

// Custom handler, implements the ServeHTTP method
type myHandler struct {
}

// Custom handler, implements the ServeHTTP method
type helloHandler struct {
}

type aboutHandler struct {
}

// ServeHTTP, not ServerHTTP
func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home handler"))
}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello handler"))
}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about handler"))
}
