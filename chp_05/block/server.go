package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if r := rand.Intn(9); r < 3 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
	} else if r < 6 && r >= 3 {
		t, _ = template.ParseFiles("layout.html", "blue_hello.html")
	} else {
		t, _ = template.ParseFiles("layout.html")
	}
	t.ExecuteTemplate(w, "layout", "I am content!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
