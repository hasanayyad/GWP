package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func daysOfTheWeek(w http.ResponseWriter, r *http.Request) {
	/*
		days := make(map[int]string)
		days[1] = "Monday"
		days[2] = "Tuesday"
		days[7] = "Sunday"
		days[6] = "Saturday"
		days[5] = "Friday"
		days[4] = "Thursday"
		days[3] = "Wednesday"
	*/

	days := map[int]string{
		1: "Monday",
		2: "Tuesday",
		7: "Sunday",
		6: "Saturday",
		5: "Friday",
		4: "Thursday",
		3: "Wednesday",
	}

	days[8] = "Eighth day of the week!"

	t, _ := template.ParseFiles("temp.html")
	t.Execute(w, days)

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/daysoftheweek", daysOfTheWeek)

	fmt.Println("Listening on port", server.Addr)
	server.ListenAndServe()
}
