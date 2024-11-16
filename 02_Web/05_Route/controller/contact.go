package controller

import (
	"fmt"
	"net/http"
	"text/template"
)

// Register routes
func registerContactRoutes() {
	http.HandleFunc("/contact", handleContact)
}

// Handle the request logic
func handleContact(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("p21/layout.html", "p21/home.html")
	if err != nil {
		fmt.Printf("Error: %v", err)                       // Ugly debug output
		w.WriteHeader(http.StatusInternalServerError) // Proper HTTP response
		return
	}
	t.ExecuteTemplate(w, "layout", "hello world")
}
