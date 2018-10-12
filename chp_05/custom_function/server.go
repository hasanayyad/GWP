package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Returns a function that "captures" the selected
// date format, using closure.
func formatDateClosure(i int) func(t time.Time) string {
	var layout string
	if i == 1 {
		layout = "2006-01-02"
	} else {
		layout = "Mon Jan 2, 2006 (3:04:05 PM)"
	}

	return func(t time.Time) string {
		return t.Format(layout)
	}
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"fdate1": formatDateClosure(1),
		"fdate2": formatDateClosure(2),
	}
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)

	fmt.Printf("Port %s is now live...\n", server.Addr)
	server.ListenAndServe()
}
