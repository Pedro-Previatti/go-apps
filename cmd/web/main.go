package main

import (
	"fmt"
	"net/http"

	"github.com/Pedro-Previatti/go-apps/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("application started \nrunning on port%s", portNumber))
	// listens to port 8080
	_ = http.ListenAndServe(portNumber, nil)
}
