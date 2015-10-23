package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	mux := http.NewServeMux()
  mux.HandleFunc("/", hello)

  n := negroni.Classic()
  n.UseHandler(mux)
	hostString := fmt.Sprintf(":%s", port)
  n.Run(hostString)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello from Go!")
}
