package main

import (
	"fmt"
	"go-web-zero/p21/controller"
	"net/http"
)

func main() {

	/* Written inside the controller
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("layout.html", "home.html")
		t.ExecuteTemplate(w, "layout", "hello world")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("layout.html", "about.html")
		t.ExecuteTemplate(w, "layout", "hello world")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		t, e := template.ParseFiles("layout.html")
		t.ExecuteTemplate(w, "layout", "hello world")
		log.Println(e)
	})
	*/
	fmt.Println("Preparing to register routes")
	controller.RegisterRoutes() // Register routes, equivalent to the code above
	fmt.Println("Routes have been registered")
	http.ListenAndServe("localhost:8080", nil) // Use the default multiplexer
}
