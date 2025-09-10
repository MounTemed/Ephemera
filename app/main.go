package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go + Ephemera + LaiMicK!!")
	})

	fmt.Println("Starting server at port 443")
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
