package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("The about returned value was %d", sum))
}

// addValues adds two values
// if it starts with a lower case it means the function is not going to leave
// the scope of this file, otherwise it can be called from anywhere
func addValues(x, y int) int {
	return x + y
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("application started \nrunning on port%s", portNumber))
	// listens to port 8080
	_ = http.ListenAndServe(portNumber, nil)
}
