package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	host := os.Getenv("HOST")
	if len(host) == 0 {
		host = "0.0.0.0"
	}

	listenOn := host + ":" + port

	fmt.Printf("Server listening on: %s\n", listenOn)
	err := http.ListenAndServe(listenOn, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello")
}
