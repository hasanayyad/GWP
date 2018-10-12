package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	// This is the message that will be URL-encoded into the flash cookie
	message := []byte("This is a flash (temporary) cookie...")

	// Make the new flash cookie
	cookie := http.Cookie{
		Name:  "flashCookie",
		Value: base64.URLEncoding.EncodeToString(message),
	}

	// Send the cookie to the browser, and notify via HTML message
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "flashCookie is ready!")
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	// Check if the flash cookie has already been set
	cookie, err := r.Cookie("flashCookie")
	if err != nil {
		// If the flash cookie hasn't been set, display a message that says so
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "Cookie not found or expired...")
		}
		
		// If the flash cookie has been set, make another cookie to displace it,
		// and give the new cookie a negative age (as well as an expiration time in the past)
		// to make sure it is expired.
	} else {
		newCookie := &http.Cookie{
			Name:    "flashCookie",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		// Send the new, expired cookie back to the browser
		http.SetCookie(w, newCookie)
		// Print the message from the old flash cookie
		message, _ := base64.URLEncoding.DecodeString(cookie.Value)
		fmt.Fprintln(w, string(message))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}
