package controller

import (
	"fmt"
	"net/http"
	"text/template"
)

// Register routes
func registerAboutRoutes() {
	http.HandleFunc("/about", handleAbout)
}

// Handle the request logic
func handleAbout(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("p21/layout.html", "p21/about.html")
	// If the server cannot find the data, check the error. 
	// Error message might be: open layout.html: The system cannot find the file specified.
	if err != nil {
		fmt.Println(err)                              // Ugly debug output
		w.WriteHeader(http.StatusInternalServerError) // Proper HTTP response, HTTP ERROR 500
		return
	}
	t.ExecuteTemplate(w, "about", "hello world")
}
