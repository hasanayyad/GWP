package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `r.FormValue("hello"):`, r.FormValue("hello"))
	fmt.Fprintln(w, `r.PostFormValue("hello"):`, r.PostFormValue("hello"))
	fmt.Fprintln(w, `r.PostForm:`, r.PostForm)
	fmt.Fprintln(w, `r.MultipartForm:`, r.MultipartForm)
	fmt.Fprintln(w, `r.Form:`, r.Form)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
