package main

import (
	// "fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	// HandleFunc(patter string, handler func(ResponseWriter, *Request))
	// HandleFunc registers the handler function for the given patter in the DefultServerMux

	// ListenAndServer(addr string, handler Handler) error
	// listens on the TCP address and calls Serve handler to HANDLE reqs

	// templates are typically strings or files containing placeholders or special syntax for dynamic content
	// Parsing is just reading the template for those strings, plaecholders, and special syntax
	// exmp ~> templateString := "Hello, {{.Name}}"

	// tpl, _ = template.ParseFiles("index.html") // , _ we aren't accounting for errors at the moment so this a placeholder throwaway for err
	// tpl, _ = template.ParseFiles("data1/index2.html") // After main.go has been searched, Go looks for the specific path to the template being read
	// tpl, _ = template.ParseFiles("data1/data2/index3.html")

	// The above i useful in showcasing how templates are wr, re
	// modern web development requires more than a single entry
	// func (t *Template) ParseGlob(pattern string) (*Template, error)
	// ParseGlob looks for a mathing pattern and Parses all files at once ~ convenient to bulk-load
	// tpl, _ = template.ParseGlob("templates/*.html") // * is the wildcard allowing us to have anything in front as long as it ends with html
	tpl, _ = tpl.ParseGlob("templates/*.html") // works the same as above

	// registered route handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandleFunc)
	http.HandleFunc("/about", aboutHandleFunc)

	http.ListenAndServe(":5555", nil) // nill invokes DefaultServeMux
	// ServeMux is an HTTP request multiplexer. It matches the URL of each incoming req
	// against a list of registered patterns and calls the handler for the pattern that closely matches
	// /hello will invoke helloHandleFunc
}

// w arg provides methods to send data back to the client if needed.
// it allows the construction of an HTTP request
// r arg contains info about the incoming HTTP req, like the method, headers, query params, and body

// func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "r.URL.Path:, %s", r.URL.Path)
// }

// func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "This is Go powered website")
// }

// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{})

// allows us read and write multiple templates at once

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Take the template and process it to produce desired HTML, CSS, etc
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "hello.html", nil)
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.html", nil)
}
