package main

import (
	"html/template"
	"net/http"
	"fmt"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	fmt.Println("Listening on port", server.Addr)
	server.ListenAndServe()
	
}
