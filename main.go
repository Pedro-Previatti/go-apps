package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("application started \nrunning on port%s", portNumber))
	// listens to port 8080
	_ = http.ListenAndServe(portNumber, nil)
}
