package main

import (
	"fmt"
	"net/http"
)

func main() {

	// when the application starts the page response should be 'Hello, world!'
	// and the console should print 13 as the number of bytes got from Fprintf
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		// error handling will print the error on console
		// not the best handling but simple for this application
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(fmt.Sprintf("Numbers of bytes written: %d", n))
	})

	// listens to port 8080
	_ = http.ListenAndServe(":8080", nil)
}
