package main

import (
	"html/template"
	"net/http"
)

// Custom handler
func process(w http.ResponseWriter, r *http.Request) {
	// Parse the template file, relative path
	t, _ := template.ParseFiles("tmpl.html")
	// Execute the template, pass data to replace {{.}}
	t.Execute(w, "hello Kitty")
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/test", process)
	server.ListenAndServe()
}
