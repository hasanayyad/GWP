package main

import (
	"fmt"
	"net/http"
	"time"
)

// Person is a person and a person.
type Person struct {
	Name  string
	Value string
	Age   int64
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie1 := http.Cookie{
		Name:     "first_cookie",
		Value:    time.Now().Format("2006/01/02 03:04:05 PM"),
		HttpOnly: true,
	}

	cookie2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Hi! I am the second cookie!",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)

	fmt.Fprintln(w, "The first and second cookies have been set!")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie1, err1 := r.Cookie("first_cookie")
	if err1 != nil {
		fmt.Fprintln(w, "Could not get the first cookie.")
	}

	cookie2, err2 := r.Cookie("second_cookie")
	if err2 != nil {
		fmt.Fprintln(w, "Could not get the second cookie.")
	}

	fmt.Fprintln(w, cookie1)
	fmt.Fprintln(w, cookie2)

	cookies := r.Cookies()
	fmt.Fprintln(w, cookies)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}

func testFunc(input int) (output int) {
	output = input * 2
	return
}
