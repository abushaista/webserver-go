package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	mux.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "test")
	})

	server := http.Server{Addr: ":8080", Handler: mux}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
