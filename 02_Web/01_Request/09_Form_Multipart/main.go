package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(1024)
	//
	// fileHeader := r.MultipartForm.File["uploaded"][0] // Get the first uploaded file
	// file, err := fileHeader.Open()

	file, _, err := r.FormFile("uploaded") // The corresponding key in the HTML form
	// This will only return the first file, suitable for single-file uploads
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}

}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
