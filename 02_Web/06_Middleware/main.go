package main

import (
	"encoding/json"
	"go-web-zero/p23/middleware"
	"net/http"
)

type Company struct {
	ID      int
	Name    string
	Country string
}

func main() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		c := Company{
			ID:      123,
			Name:    "Google",
			Country: "USA",
		}
		enc := json.NewEncoder(w)
		enc.Encode(c)
	})
	// use middleware
	http.ListenAndServe("localhost:8080", new(middleware.AuthMiddleware))
}
