package main

import (
	"encoding/json"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Go web</title>
</head>
<body>
    hello world
</body>
</html>`

	w.Write([]byte(str)) // Write str to the body, requires type conversion
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://apple.com")
	// The header must be set before writing the response
	w.WriteHeader(302) // Redirect with status code 302
}

type POST struct {
	User   string
	Thread []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &POST{
		User:   "lord",
		Thread: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {

	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/redirect", headerExample)

	// JSON example
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()

}
