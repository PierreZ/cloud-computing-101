package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	if message == "bug" {
		fmt.Println("oups")
		os.Exit(1)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	message = "Hello '" + message + "' from " + hostname

	w.Write([]byte(message))
}

func main() {

	port := os.Getenv("MY_APP_LISTEN_PORT")
	if port == "" {
		port = "8080"
	}

	parsedPort, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	listenArg := fmt.Sprintf(":%d", parsedPort)

	http.HandleFunc("/", sayHello)
	fmt.Println("starting server on " + listenArg)
	if err := http.ListenAndServe(listenArg, nil); err != nil {
		panic(err)
	}
}
