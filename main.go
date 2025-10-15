package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Parse all templates
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		tmpl.ExecuteTemplate(w, "index.html", map[string]string{
			"Title": "Home Page",
			"Body":  "Welcome to my first Go website 🚀",
		})
	case "/about":
		tmpl.ExecuteTemplate(w, "about.html", map[string]string{
			"Title": "About Page",
			"Body":  "Main brach!",
		})
	default:
		http.NotFound(w, r)
	}
}

func main() {
	// Serve static files (CSS, JS, images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle routes
	http.HandleFunc("/", handler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
