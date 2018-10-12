package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func header(w http.ResponseWriter, r *http.Request) {
	h := r.Header[r.URL.Path]
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	http.Handle("/headers/", http.StripPrefix("/headers/", http.HandlerFunc(header)))
	server.ListenAndServe()
}
