package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	if message == "bug" {
		fmt.Println("oups")
		os.Exit(1)
	}
	message = "Hello " + message

	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	fmt.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
