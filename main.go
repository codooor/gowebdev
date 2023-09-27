package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	tpl, _ = template.ParseFiles("index.html")
	// handle http req's at the path /hello
	// ListenandServe starts the server HTTP serveer and listens for any req's on it's path
	http.HandleFunc("/hello", helloHandleFunc) // helloHandleFunc is registered the /hello path
	http.HandleFunc("/about", aboutHandleFunc)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":5555", nil) // nill invokes DefaultServeMux
	// ServeMux is an HTTP request multiplexer. It matches the URL of each incoming req
	// against a list of registered patterns and calls the handler for the pattern that closely matches
	// /hello will invoke helloHandleFunc
}

// w arg provides methods to send data back to the client if needed.
// it allows the construction of an HTTP request
// r arg contains info about the incoming HTTP req, like the method, headers, query params, and body
func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "r.URL.Path:, %s", r.URL.Path)
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is Go powered website")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
