package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

// renderTemplate is a function to open and render go templates
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("application started \nrunning on port%s", portNumber))
	// listens to port 8080
	_ = http.ListenAndServe(portNumber, nil)
}
