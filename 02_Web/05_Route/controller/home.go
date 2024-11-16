package controller

import (
	"log"
	"net/http"
	"text/template"
)

// Register routes
func registerHomeRoutes() {
	http.HandleFunc("/home", handleHome)
}

// Handle the request logic
func handleHome(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("p21/layout.html", "p21/home.html")
	log.Println(e)
	t.ExecuteTemplate(w, "layout", "hello world")
}
